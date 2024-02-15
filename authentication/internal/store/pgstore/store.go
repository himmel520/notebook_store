package pgstore

import (
	"database/sql"

	"github.com/himmel520/notebook_store/authentication/internal/store"

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
