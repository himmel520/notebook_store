package server

import (
	"encoding/json"
	"net/http"
	log "store/internal/logger"
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

	log.Logger.Info("start server")
	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/catalog", s.handleGetAllNotebookInfo()).Methods("GET")
	s.router.HandleFunc("/catalog/{id:[0-9]+}", s.handleGetNotebookInfo()).Methods("GET")

	admin := s.router.NewRoute().Subrouter()
	admin.Use(s.handleAuthorization)
	admin.HandleFunc("/addSystem", s.handleAddSystem()).Methods("POST")
	admin.HandleFunc("/addScreen", s.handleAddScreen()).Methods("POST")
	admin.HandleFunc("/addProcessor", s.handleAddProcessor()).Methods("POST")
	admin.HandleFunc("/addStorage", s.handleAddStorage()).Methods("POST")
	admin.HandleFunc("/addRam", s.handleAddRam()).Methods("POST")
	admin.HandleFunc("/addNotebook", s.handleAddNotebook()).Methods("POST")
	admin.HandleFunc("/deleteNotebook/{id:[0-9]+}", s.handleDeleteNotebook()).Methods("DELETE")
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
			log.Logger.Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleGetAllNotebookInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notebooks, err := s.store.Notebook().GetAllNotebooks()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(notebooks)
		if err != nil {
			log.Logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (s *Server) handleGetNotebookInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		notebookInfo := &models.NotebookInfo{}
		if err := s.store.Notebook().FindNotebookByID(id, notebookInfo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jsonData, err := json.Marshal(notebookInfo)
		if err != nil {
			log.Logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
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
		processor := &models.Processor{}
		if err := json.NewDecoder(r.Body).Decode(processor); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Components().CreateProcessor(processor); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	}
}

func (s *Server) handleAddStorage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		storage := &models.Storage{}
		if err := json.NewDecoder(r.Body).Decode(storage); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Components().CreateStorage(storage); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (s *Server) handleAddRam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ram := &models.RAM{}
		if err := json.NewDecoder(r.Body).Decode(ram); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Components().CreateRam(ram); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

	}
}

func (s *Server) handleAddNotebook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notebook := &models.Notebook{}
		if err := json.NewDecoder(r.Body).Decode(notebook); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := s.store.Notebook().CreateNotebook(notebook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (s *Server) handleDeleteNotebook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if err := s.store.Notebook().DeleteNotebookByID(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
