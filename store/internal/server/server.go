package server

import (
	"encoding/json"
	"net/http"
	"store/internal/models"
	"store/internal/store"
	"time"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
)

type Server struct {
	config *Config
	nats   *nats.Conn
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

	db, err := s.NewDB(s.config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	s.ConfigureStore(db)
	s.ConfigureRouter()

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/catalog", nil).Methods("GET")
	s.router.HandleFunc("/catalog/id", nil).Methods("GET")

	admin := s.router.NewRoute().Subrouter()
	admin.Use(s.handleAuthorization)
	admin.HandleFunc("/addSystem", s.handleAddSystem()).Methods("POST")
	admin.HandleFunc("/addScreen", s.handleAddScreen()).Methods("POST")
	admin.HandleFunc("/addProcessor", nil).Methods("POST")
	admin.HandleFunc("/addStorage", nil).Methods("POST")
	admin.HandleFunc("/addRam", nil).Methods("POST")
	admin.HandleFunc("/addNotebook", nil).Methods("POST")
	admin.HandleFunc("/deleteNotebook", nil).Methods("DELETE")
	admin.HandleFunc("/getAllData", nil).Methods("GET")
}

func (s *Server) handleAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		response, err := s.nats.Request("authorization", []byte(sessionId.Value), 2*time.Second)
		if err != nil || string(response.Data) != "true" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleAddSystem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		system := &models.System{}
		if err := json.NewDecoder(r.Body).Decode(system); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Components().CreateSystem(system); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (s *Server) handleAddScreen() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		screen := &models.Screen{}
		if err := json.NewDecoder(r.Body).Decode(screen); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Components().CreateScreen(screen); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (s *Server) handleAddProcessor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleAddStorage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleAddRam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleAddNotebook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleDeleteNotebook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleGetAllData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
