package server

import (
	"database/sql"

	"github.com/himmel520/notebook_store/store/internal/store/pgstore"
)

func (s *Server) NewDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Server) ConfigureStore(db *sql.DB) {
	st := pgstore.New(db)
	s.store = st
}
