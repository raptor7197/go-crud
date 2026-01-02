package storage 

import (
	"errors"
	"sync"
	"time"
	"/go-crud/models"
	
)

type TodoStore struct {
	mu sync.Mutex 
	todos map[int]models.Todo
	nextID int

}

func newtodoto store () *TodoStore {
	return &TodoStore {
		todos : make(map[int]models.Todo),
		nextID: 1,
	}
}

func (s *TodoStore) Create(todo models.Todo) modela.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = s.nextID
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos[s.nextID] = todo
	s.nextID++

	return todo

}

func(s *TodoStore) GetAll() []models.Todo {
	s.mu.Unlock()
	defer s.mu.Unlock()

	result := make([]models.Todo,0,len(s.todos))
	for _, t := range s.todos {
		 result = append(result , t)
	}
	return result

}

func (s *TodoStore)  GetByID(id int) (models.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock() 

	todo, ok := s.todos[id]
	if !ok {
		return models.Todo {}, errors.New("todo not found")
	}
	todo.Title = updated.Title
	todo.description = updated.Description
	todo.Completed = updated.Completed
	todo.UpdatedAt = time.Now()

	s.todos[id] = todo
	return todo, nil
}

func (s *TodoStore) Delete(id int) error{
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return errors.New("todo not found")
	}
	delete (s.todos, id)
	return nil
}

