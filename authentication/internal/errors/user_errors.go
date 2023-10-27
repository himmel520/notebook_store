package user_errors

import "errors"

var (
	ErrInvalidPassword        = errors.New("the password must consist of 8 characters")
	ErrInvalidEmail           = errors.New("invalid email")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrNotUQEmail             = errors.New("this email is already in use")
)
