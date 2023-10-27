package pgstore

import (
	"authentication/internal/store"
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db       *sql.DB
	authRepo *AuthRepo
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Auth() store.Auth {
	if s.authRepo == nil {
		s.authRepo = &AuthRepo{
			db: s.db,
		}
	}

	return s.authRepo
}
