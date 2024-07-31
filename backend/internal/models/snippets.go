package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content string, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, now(), now() + $3::interval) RETURNING id`

	result := m.DB.QueryRow(context.Background(), stmt, title, content, expires)

	var id pgtype.Int4
	err := result.Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id.Int32), nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > now() AND id = $1`

	var s Snippet
	err := m.DB.QueryRow(context.Background(), stmt, id).
		Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []Snippet{}

	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
