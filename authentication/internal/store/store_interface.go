package store

import "authentication/internal/models"

type Auth interface {
	CreateUser(u *models.User) error
	FindUserByEmail(loginUser *models.User) (*models.User, error)
}

type Store interface {
	Auth() Auth
}
