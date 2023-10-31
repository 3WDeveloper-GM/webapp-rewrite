package models

import (
	"database/sql"
	"time"
)

type Users struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte    //wtf, we hasing now
	Created        time.Time //this is weird
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(name, email, password string) error {
	return nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return true, nil
}
