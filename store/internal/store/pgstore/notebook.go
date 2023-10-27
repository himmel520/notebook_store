package pgstore

import (
	"database/sql"
	"store/internal/models"
)

type NotebookRepo struct {
	db *sql.DB
}

func (r *NotebookRepo) CreateNotebook(n *models.Notebook) error {
	return nil
}

func (r *NotebookRepo) DeleteNotebookByID(id int) error {
	return nil
}

func (r *NotebookRepo) FindNotebookByID(id int) (*models.Notebook, error) {
	return nil, nil
}
