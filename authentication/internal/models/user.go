package models

import (
	"crypto/sha256"
	"fmt"
	"regexp"

	uerr "github.com/himmel520/notebook_store/authentication/internal/errors"
)

const (
	salt       = "ewzxrcvghbjnkmmnbjkhg"
	validEmail = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
)

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
	IsAdmin      bool   `json:"-"`
}

func (u *User) validate() error {
	if len(u.Password) != 8 {
		return uerr.ErrInvalidPassword
	}

	if match, _ := regexp.MatchString(validEmail, u.Email); !match || len(u.Email) > 25 {
		return uerr.ErrInvalidEmail
	}

	return nil
}

func (u *User) cleanPassword() {
	u.Password = ""
}

func (u *User) CheckBeforeDB() error {
	if err := u.validate(); err != nil {
		return err
	}
	u.CreatePasswordHash()
	return nil
}

func (u *User) CreatePasswordHash() {
	hash := sha256.Sum256([]byte(u.Password + salt))
	u.PasswordHash = fmt.Sprintf("%x", hash)
	u.cleanPassword()
}

func (u *User) CompareHashPassword(secondUser *User) error {
	if !(u.PasswordHash == secondUser.PasswordHash) {
		return uerr.ErrInvalidEmailOrPassword
	}
	return nil
}
