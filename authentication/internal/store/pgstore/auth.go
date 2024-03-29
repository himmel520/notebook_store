package pgstore

import (
	"database/sql"

	uerr "github.com/himmel520/notebook_store/authentication/internal/errors"
	"github.com/himmel520/notebook_store/authentication/internal/models"
)

type AuthRepo struct {
	db *sql.DB
}

func (r *AuthRepo) CreateUser(u *models.User) error {
	if err := u.CheckBeforeDB(); err != nil {
		return err
	}

	_, err := r.db.Exec(
		`insert into users 
			(email, password_hash) 
		 values 
			($1, $2)`,
		u.Email, u.PasswordHash)
	if err != nil {
		return uerr.ErrNotUQEmail
	}
	return nil
}

func (r *AuthRepo) FindUserByEmail(loginUser *models.User) (*models.User, error) {
	if err := loginUser.CheckBeforeDB(); err != nil {
		return nil, err
	}

	var registeredUser models.User
	err := r.db.QueryRow(
		`select 
			id_users, email, password_hash, is_admin
		 from users
		 	where email=$1`, loginUser.Email).Scan(&registeredUser.ID, &registeredUser.Email, &registeredUser.PasswordHash, &registeredUser.IsAdmin)
	if err != nil {
		return nil, uerr.ErrInvalidEmailOrPassword
	}
	return &registeredUser, nil
}
