package storage

import (
	"errors"
	"go-crud/models"
	"sync"
	"time"
)

// TodoStore is a simple in-memory storage for Todos
type TodoStore struct {
	mu     sync.Mutex
	todos  map[int]models.Todo
	nextID int
}

// NewTodoStore creates a new TodoStore
func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

// Create adds a new Todo to the store
func (s *TodoStore) Create(todo models.Todo) models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos[s.nextID] = todo
	s.nextID++

	return todo
}

// GetAll returns all Todos from the store
func (s *TodoStore) GetAll() []models.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]models.Todo, 0, len(s.todos))
	for _, t := range s.todos {
		result = append(result, t)
	}
	return result
}

// GetByID returns a single Todo from the store by ID
func (s *TodoStore) GetByID(id int) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]
	if !ok {
		return models.Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

// Update updates a Todo in the store
func (s *TodoStore) Update(id int, updated models.Todo) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]
	if !ok {
		return models.Todo{}, errors.New("todo not found")
	}

	todo.Title = updated.Title
	todo.Description = updated.Description
	todo.Completed = updated.Completed
	todo.UpdatedAt = time.Now()

	s.todos[id] = todo
	return todo, nil
}

// Delete removes a Todo from the store by ID
func (s *TodoStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return errors.New("todo not found")
	}
	delete(s.todos, id)
	return nil
}

