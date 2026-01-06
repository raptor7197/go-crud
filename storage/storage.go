
package storage

import (
	"errors"
	"go-crud/models"
	"sync"
	"time"
)

// TodoStore defines the interface for todo storage
type TodoStore interface {
	Create(todo models.Todo) (models.Todo, error)
	GetAll() ([]models.Todo, error)
	GetByID(id int) (models.Todo, error)
	Update(id int, updated models.Todo) (models.Todo, error)
	Delete(id int) error
}

// InMemoryTodoStore is a simple in-memory storage for Todos
type InMemoryTodoStore struct {
	mu     sync.Mutex
	todos  map[int]models.Todo
	nextID int
}

// NewInMemoryTodoStore creates a new InMemoryTodoStore
func NewInMemoryTodoStore() *InMemoryTodoStore {
	return &InMemoryTodoStore{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

// Create adds a new Todo to the store
func (s *InMemoryTodoStore) Create(todo models.Todo) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos[s.nextID] = todo
	s.nextID++

	return todo, nil
}

// GetAll returns all Todos from the store
func (s *InMemoryTodoStore) GetAll() ([]models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]models.Todo, 0, len(s.todos))
	for _, t := range s.todos {
		result = append(result, t)
	}
	return result, nil
}

// GetByID returns a single Todo from the store by ID
func (s *InMemoryTodoStore) GetByID(id int) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[id]
	if !ok {
		return models.Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

// Update updates a Todo in the store
func (s *InMemoryTodoStore) Update(id int, updated models.Todo) (models.Todo, error) {
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
func (s *InMemoryTodoStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return errors.New("todo not found")
	}
	delete(s.todos, id)
	return nil
}

