package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	SQL_STATEMENT := `INSERT INTO users (name, email, hashed_password, created) VALUES (?,?,?,UTC_TIMESTAMP());`

	_, err = m.DB.Exec(SQL_STATEMENT, name, email, string(HashedPassword))
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (m *UserModel) Authenticate(name, email, password string) error {
	return nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return true, nil
}
