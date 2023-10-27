package pgstore

import (
	"database/sql"
	merr "store/internal/errors"
	"store/internal/models"
)

type ComponentsRepo struct {
	db *sql.DB
}

func (r *ComponentsRepo) CreateSystem(s *models.System) error {
	if err := s.Validate(); err != nil {
		return err
	}

	_, err := r.db.Exec(`insert into systems (name) values ($1)`, s.Name)
	if err != nil {
		return merr.ErrNotUQSystem
	}

	return nil
}

func (r *ComponentsRepo) CreateScreen(s *models.Screen) error {
	if err := s.Validate(); err != nil {
		return err
	}

	_, err := r.db.Exec(
		`insert into screens (size_inches, resolution) values ($1, $2)`,
		s.SizeInInches, s.Resolution)
	if err != nil {
		return merr.ErrorOccurred
	}

	return nil
}

func (r *ComponentsRepo) CreateProcessor(p *models.Processor) error {
	return nil
}

func (r *ComponentsRepo) CreateStorage(s *models.Storage) error {
	return nil
}

func (r *ComponentsRepo) CreateRam(ram *models.RAM) error {
	return nil
}
