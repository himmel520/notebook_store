package models

import merr "store/internal/errors"

type Notebook struct {
	ID          int     `json:"id"`
	SystemID    int     `json:"system_id"`
	ScreenID    int     `json:"screen_id"`
	ProcessorID int     `json:"processor_id"`
	StorageID   int     `json:"storage_id"`
	RAMID       int     `json:"ram_id"`
	Model       string  `json:"model"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type System struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *System) Validate() error {
	if len(s.Name) > 30 {
		return merr.ErrInvalidSystemName
	}
	return nil
}

type Screen struct {
	ID           int     `json:"id"`
	SizeInInches float32 `json:"size_inches"`
	Resolution   string  `json:"resolution"`
}

func (s *Screen) Validate() error {
	if s.SizeInInches > 100 || s.SizeInInches < 0 {
		return merr.ErrInvalidScreenSize
	}

	if len(s.Resolution) > 9 {
		return merr.ErrInvalidScreenResolution
	}

	return nil
}

type Processor struct {
	ID         int     `json:"id"`
	Model      string  `json:"model"`
	SpeedInGHz float32 `json:"speed_GHz"`
}

func (p *Processor) Validate() error {
	return nil
}

type Storage struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	SizeInGB int    `json:"size_gb"`
}

func (s *Storage) Validate() error {
	return nil
}

type RAM struct {
	ID       int `json:"id"`
	SizeInGB int `json:"size_gb"`
}

func (r *RAM) Validate() error {
	return nil
}
