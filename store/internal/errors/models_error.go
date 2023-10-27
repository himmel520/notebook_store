package models_errors

import "errors"

var (
	ErrInvalidSystemName       = errors.New("the system name must not exceed 30 characters")
	ErrInvalidScreenSize       = errors.New("the number should be positive and less than 100")
	ErrInvalidScreenResolution = errors.New("the screen resolutin must not exceed 9 characters")
)

var (
	ErrNotUQSystem = errors.New("this system name is already in use")
	ErrorOccurred  = errors.New("something went wrong")
)
