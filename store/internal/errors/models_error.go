package models_errors

import "errors"

var (
	ErrInvalidSystemName        = errors.New("the system name must not exceed 30 characters")
	ErrInvalidScreenSize        = errors.New("the screen size should be positive and less than 100")
	ErrInvalidScreenResolution  = errors.New("the screen resolutin must not exceed 9 characters")
	ErrInvalidProcessorModel    = errors.New("the model must not exceed 30 characters")
	ErrInvalidProcessorSpeedGHZ = errors.New("the speed(ghz) should be positive and less than 10")
	ErrInvalidSizeGB            = errors.New("the size(GB) should be positive")
	ErrInvalidStorageType       = errors.New("the type must not exceed 10 characters")
	ErrInvalidNotebookModel     = errors.New("the notebook model must not exceed 30 characters")
	ErrInvalidNotebookPrice     = errors.New("the notebook price should be positive and less than 10 000 000")
)

var (
	ErrNotUQSystem = errors.New("this system name is already in use")
	ErrorOccurred  = errors.New("something went wrong")
)
