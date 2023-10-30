package pgstore

import (
	"database/sql"
	merr "store/internal/errors"
	log "store/internal/logger"
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
		log.Logger.Warning(err)
		return merr.ErrorOccurred
	}

	return nil
}

func (r *ComponentsRepo) CreateProcessor(p *models.Processor) error {
	if err := p.Validate(); err != nil {
		return err
	}

	_, err := r.db.Exec(`insert into processors (model, speed_ghz) values ($1, $2)`,
		p.Model, p.SpeedInGHz)
	if err != nil {
		log.Logger.Warning(err)
		return merr.ErrorOccurred
	}
	return nil

}

func (r *ComponentsRepo) CreateStorage(s *models.Storage) error {
	if err := s.Validate(); err != nil {
		return err
	}

	_, err := r.db.Exec(`insert into storages (type_storage, size_gb) values ($1, $2)`,
		s.Type, s.SizeInGB)
	if err != nil {
		log.Logger.Warning(err)
		return merr.ErrorOccurred
	}

	return nil
}

func (r *ComponentsRepo) CreateRam(ram *models.RAM) error {
	if err := ram.Validate(); err != nil {
		return err
	}

	_, err := r.db.Exec(`insert into rams (size_gb) values ($1)`, ram.SizeInGB)
	if err != nil {
		log.Logger.Warning(err)
		return merr.ErrorOccurred
	}

	return nil
}
