package models

import merr "github.com/himmel520/notebook_store/store/internal/errors"

type NotebookInfo struct {
	SystemName          string  `json:"system_name"`
	ScreenSizeInches    float64 `json:"screen_size_inches"`
	ScreenResolution    string  `json:"screen_resolution"`
	ProcessorModel      string  `json:"processor_model"`
	ProcessorSpeedGHz   float64 `json:"processor_speed_ghz"`
	StorageType         string  `json:"storage_type"`
	StorageSizeGB       int     `json:"storage_size_gb"`
	RAMSizeGB           int     `json:"ram_size_gb"`
	NotebookModel       string  `json:"notebook_model"`
	NotebookDescription string  `json:"notebook_description"`
	NotebookPrice       float64 `json:"notebook_price"`
}

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

func (n *Notebook) Validate() error {
	if len(n.Model) > 30 {
		return merr.ErrInvalidNotebookModel
	}

	if n.Price >= 10_000_000 || n.Price < 1 {
		return merr.ErrInvalidNotebookPrice
	}

	return nil
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
	if s.SizeInInches > 100 || s.SizeInInches < 1 {
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
	if len(p.Model) > 30 {
		return merr.ErrInvalidProcessorModel
	}

	if p.SpeedInGHz > 10 || p.SpeedInGHz < 1 {
		return merr.ErrInvalidProcessorSpeedGHZ
	}

	return nil
}

type Storage struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	SizeInGB int    `json:"size_gb"`
}

func (s *Storage) Validate() error {
	if s.SizeInGB < 1 {
		return merr.ErrInvalidSizeGB
	}

	if len(s.Type) > 10 {
		return merr.ErrInvalidStorageType
	}

	return nil
}

type RAM struct {
	ID       int `json:"id"`
	SizeInGB int `json:"size_gb"`
}

func (r *RAM) Validate() error {
	if r.SizeInGB < 1 {
		return merr.ErrInvalidSizeGB
	}
	return nil
}
