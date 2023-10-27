package server

import (
	"authentication/internal/store/pgstore"
	"database/sql"
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
