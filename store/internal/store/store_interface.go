package store

import "github.com/himmel520/notebook_store/store/internal/models"

type Notebook interface {
	CreateNotebook(n *models.Notebook) error
	DeleteNotebookByID(id string) error
	FindNotebookByID(id string, info *models.NotebookInfo) error
	GetAllNotebooks() ([]*models.NotebookInfo, error)
}

type Components interface {
	CreateSystem(s *models.System) error
	CreateScreen(s *models.Screen) error
	CreateProcessor(p *models.Processor) error
	CreateStorage(s *models.Storage) error
	CreateRam(ram *models.RAM) error
}

type Store interface {
	Notebook() Notebook
	Components() Components
}
