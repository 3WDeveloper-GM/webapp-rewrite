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

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	SQL_STATEMENT := `SELECT id, hashed_password FROM users WHERE email = ?`

	err := m.DB.QueryRow(SQL_STATEMENT, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Checks if the row exist
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	var Exists bool

	SQL_STATEMENT := "SELECT EXISTS(SELECT true FROM users WHERE id = ?);"

	err := m.DB.QueryRow(SQL_STATEMENT, id).Scan(&Exists)
	return Exists, err
}

func (m *UserModel) Get(id int) (*Users, error) {
	var UserData *Users

	SQL_STATEMENT := `SELECT name, email, created FROM users WHERE id =?;`

	err := m.DB.QueryRow(SQL_STATEMENT, id).Scan(&UserData.Name, &UserData.Email, &UserData.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return UserData, err
}
