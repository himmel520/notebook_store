package pgstore

import (
	"database/sql"

	"github.com/himmel520/notebook_store/store/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	notebookRepo   *NotebookRepo
	componentsRepo *ComponentsRepo
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Notebook() store.Notebook {
	if s.notebookRepo == nil {
		s.notebookRepo = &NotebookRepo{
			db: s.db,
		}
	}

	return s.notebookRepo
}

func (s *Store) Components() store.Components {
	if s.componentsRepo == nil {
		s.componentsRepo = &ComponentsRepo{
			db: s.db,
		}
	}

	return s.componentsRepo
}
