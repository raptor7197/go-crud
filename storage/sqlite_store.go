package storage

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"go-crud/models"
	"time"
)

type SqliteTodoStore struct {
	db *sql.DB
}

func NewSqliteTodoStore(dsn string) (*SqliteTodoStore, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	store := &SqliteTodoStore{db: db}
	err = store.init()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *SqliteTodoStore) init() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			completed BOOLEAN NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL)`)
	return err
}

func (s *SqliteTodoStore) Create(todo models.Todo) (models.Todo, error) {
	now := time.Now()

	res, err := s.db.Exec(`INSERT INTO todos (title, description, completed, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`, todo.Title, todo.Description, todo.Completed, now, now)
	if err != nil {
		return models.Todo{}, err
	}
	id, _ := res.LastInsertId()
	todo.ID = int(id)
	todo.CreatedAt = now
	todo.UpdatedAt = now

	return todo, nil
}

func (s *SqliteTodoStore) GetAll() ([]models.Todo, error) {
	rows, err := s.db.Query(`SELECT id, title, description, completed, created_at, updated_at FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (s *SqliteTodoStore) GetByID(id int) (models.Todo, error) {
	var t models.Todo

	err := s.db.QueryRow(`
		SELECT id, title, description, completed, created_at, updated_at
		FROM todos WHERE id = ?
	`, id).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)

	if err == sql.ErrNoRows {
		return models.Todo{}, errors.New("todo not found")
	}

	return t, err
}

func (s *SqliteTodoStore) Update(id int, updated models.Todo) (models.Todo, error) {
	now := time.Now()
	res, err := s.db.Exec(`UPDATE todos
		SET title = ?, description = ?, completed = ?, updated_at = ?
		WHERE id = ?
	`, updated.Title, updated.Description, updated.Completed, now, id)
	if err != nil {
		return models.Todo{}, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return models.Todo{}, errors.New("todo not found")
	}
	return s.GetByID(id)
}

func (s *SqliteTodoStore) Delete(id int) error {
	res, err := s.db.Exec(`DELETE FROM todos WHERE id=?`, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("todo not found")
	}
	return nil
}
