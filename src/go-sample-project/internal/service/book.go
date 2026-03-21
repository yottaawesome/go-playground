// Package service contains the business logic of the application.
// It is decoupled from HTTP concerns and could be reused with gRPC, CLI, etc.
package service

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/example/go-sample-project/internal/model"
)

// Sentinel errors returned by the service layer.
// Using errors.New makes it easy for callers to check with errors.Is.
var (
	ErrBookNotFound = errors.New("book not found")
	ErrInvalidInput = errors.New("invalid input: title and author are required")
)

// BookService manages the collection of books.
// In a real application this would talk to a database; here we use an
// in-memory map protected by a read-write mutex for simplicity.
type BookService struct {
	mu    sync.RWMutex
	books map[string]model.Book
}

// NewBookService creates an initialised BookService.
func NewBookService() *BookService {
	return &BookService{
		books: make(map[string]model.Book),
	}
}

// List returns all books in no guaranteed order.
func (s *BookService) List() []model.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Pre-allocate the slice to avoid repeated growing.
	result := make([]model.Book, 0, len(s.books))
	for _, b := range s.books {
		result = append(result, b)
	}
	return result
}

// GetByID retrieves a single book by its ID.
func (s *BookService) GetByID(id string) (model.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, ok := s.books[id]
	if !ok {
		return model.Book{}, ErrBookNotFound
	}
	return book, nil
}

// Create validates the request and persists a new book.
func (s *BookService) Create(req model.CreateBookRequest) (model.Book, error) {
	if req.Title == "" || req.Author == "" {
		return model.Book{}, ErrInvalidInput
	}

	now := time.Now().UTC()
	book := model.Book{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Author:    req.Author,
		ISBN:      req.ISBN,
		Pages:     req.Pages,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.mu.Lock()
	s.books[book.ID] = book
	s.mu.Unlock()

	return book, nil
}

// Update modifies an existing book. All provided fields overwrite the old values.
func (s *BookService) Update(id string, req model.UpdateBookRequest) (model.Book, error) {
	if req.Title == "" || req.Author == "" {
		return model.Book{}, ErrInvalidInput
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	book, ok := s.books[id]
	if !ok {
		return model.Book{}, ErrBookNotFound
	}

	// Apply changes while preserving the original creation timestamp.
	book.Title = req.Title
	book.Author = req.Author
	book.ISBN = req.ISBN
	book.Pages = req.Pages
	book.UpdatedAt = time.Now().UTC()

	s.books[id] = book
	return book, nil
}

// Delete removes a book by ID. Returns an error if the book does not exist.
func (s *BookService) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.books[id]; !ok {
		return ErrBookNotFound
	}
	delete(s.books, id)
	return nil
}
