package todo

import (
	"context"
	"errors"
)

// Predefined errors.
// HINT: Define domain-specific errors here.
var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrInvalidInput = errors.New("invalid input")
)

// Inputs for service methods (decouple from HTTP JSON).
type CreateTodoInput struct {
	Title string
}

// Inputs for service methods (decouple from HTTP JSON).
type UpdateTodoInput struct {
	Title     string
	Completed bool
}

// Service defines the business logic for managing todos.
type Service interface {
	ListTodos(ctx context.Context) ([]Todo, error)
	GetTodo(ctx context.Context, id uint) (*Todo, error)
	CreateTodo(ctx context.Context, in CreateTodoInput) (*Todo, error)
	UpdateTodo(ctx context.Context, id uint, in UpdateTodoInput) (*Todo, error)
	DeleteTodo(ctx context.Context, id uint) error
}

// service is the implementation of the Service interface.
// HINT: This is how we 'inject' the repository to the service struct
type service struct {
	repo Repository
}

// NewService creates a new Service with the given Repository.
// HINT: A 'constructor' function for the Service struct.
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// ListTodos retrieves all todo items.
// HINT: Services accept 'context.Context'
func (s *service) ListTodos(ctx context.Context) ([]Todo, error) {
	// ctx is here if later you want to pass it to repo/db
	return s.repo.FindAll()
}

// GetTodo retrieves a todo item by its ID.
func (s *service) GetTodo(ctx context.Context, id uint) (*Todo, error) {
	todo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, ErrTodoNotFound
	}
	return todo, nil
}

// CreateTodo creates a new todo item.
func (s *service) CreateTodo(ctx context.Context, todoInput CreateTodoInput) (*Todo, error) {
	if todoInput.Title == "" {
		return nil, ErrInvalidInput
	}

	todo := &Todo{
		Title: todoInput.Title,
	}
	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// UpdateTodo updates an existing todo item.
func (s *service) UpdateTodo(ctx context.Context, id uint, in UpdateTodoInput) (*Todo, error) {
	if in.Title == "" {
		return nil, ErrInvalidInput
	}

	todo, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, ErrTodoNotFound
	}

	todo.Title = in.Title
	todo.Completed = in.Completed

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// DeleteTodo deletes a todo item by its ID.
func (s *service) DeleteTodo(ctx context.Context, id uint) error {
	todo, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if todo == nil {
		return ErrTodoNotFound
	}
	return s.repo.Delete(id)
}
