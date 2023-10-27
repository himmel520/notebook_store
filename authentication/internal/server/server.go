package server

import (
	"authentication/internal/models"
	"authentication/internal/store"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
)

type Server struct {
	config *Config
	nats   *nats.Conn
	redis  *Redis
	store  store.Store
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	nc, err := nats.Connect(s.config.NatsURL)
	if err != nil {
		return err
	}
	defer nc.Close()
	s.nats = nc

	redis, err := NewRedis(s.config.Redis)
	if err != nil {
		return err
	}
	s.redis = redis

	db, err := s.NewDB(s.config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	s.ConfigureNatsRouter()
	s.ConfigureStore(db)
	s.ConfigureRouter()

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *Server) ConfigureNatsRouter() {
	s.nats.Subscribe("authorization", func(msg *nats.Msg) {
		role, err := s.redis.GetIsAdmin(string(msg.Data))
		if err != nil {
			return
		}
		s.nats.Publish(msg.Reply, []byte(role))
	})
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/signUp", s.handleSignUp()).Methods("POST")
	s.router.HandleFunc("/logIn", s.handleLogIn()).Methods("POST")
	s.router.HandleFunc("/logOut", s.logOut()).Methods("POST")
}

func (s *Server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
		if err := json.NewDecoder(r.Body).Decode(u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Auth().CreateUser(u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (s *Server) handleLogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginUser := &models.User{}
		if err := json.NewDecoder(r.Body).Decode(loginUser); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		registeredUser, err := s.store.Auth().FindUserByEmail(loginUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if err := loginUser.CompareHashPassword(registeredUser); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sessionId := uuid.New().String()
		if err := s.redis.SetSessionID(sessionId, registeredUser); err != nil {
			log.Println("[REDIS]", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionId,
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour * 30),
		})
	}
}

func (s *Server) logOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := s.redis.DeleteSessionID(sessionId.Value); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
