package store

import "store/internal/models"

type Notebook interface {
	CreateNotebook(n *models.Notebook) error
	DeleteNotebookByID(id int) error
	FindNotebookByID(id int) (*models.Notebook, error)
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
