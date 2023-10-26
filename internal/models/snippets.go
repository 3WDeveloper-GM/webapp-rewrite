package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type Snippetmodel struct {
	DB *sql.DB
}

// Inserts a snippet into the database
func (m *Snippetmodel) Insert(title string, content string, expires int) (int, error) {
	//Inserting the SQL statement
	SQL_statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY));`

	result, err := m.DB.Exec(SQL_statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *Snippetmodel) Get(id int) (*Snippet, error) {

	SQL_statement := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?;`

	row := m.DB.QueryRow(SQL_statement, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *Snippetmodel) Latest() ([]*Snippet, error) {
	return nil, nil
}
