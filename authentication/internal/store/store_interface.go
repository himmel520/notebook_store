package store

import "github.com/himmel520/notebook_store/authentication/internal/models"

type Auth interface {
	CreateUser(u *models.User) error
	FindUserByEmail(loginUser *models.User) (*models.User, error)
}

type Store interface {
	Auth() Auth
}
