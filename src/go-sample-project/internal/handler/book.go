// Package handler contains the HTTP handlers that translate between
// incoming HTTP requests and the service layer. Each handler is responsible
// for parsing input, calling the service, and writing the JSON response.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/example/go-sample-project/internal/model"
	"github.com/example/go-sample-project/internal/service"
	"github.com/example/go-sample-project/pkg/response"
)

// BookHandler holds a reference to the BookService.
// Using an interface here would make it easy to swap in a mock for testing.
type BookHandler struct {
	svc *service.BookService
}

// NewBookHandler returns a handler wired to the given service.
func NewBookHandler(svc *service.BookService) *BookHandler {
	return &BookHandler{svc: svc}
}

// ListBooks responds with the full list of books (GET /api/v1/books).
func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books := h.svc.List()
	response.JSON(w, http.StatusOK, books)
}

// GetBook responds with a single book by ID (GET /api/v1/books/{id}).
func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := h.svc.GetByID(id)
	if errors.Is(err, service.ErrBookNotFound) {
		response.Error(w, http.StatusNotFound, "book not found")
		return
	}

	response.JSON(w, http.StatusOK, book)
}

// CreateBook parses a JSON body and creates a new book (POST /api/v1/books).
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Warn().Err(err).Msg("invalid request body")
		response.Error(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	book, err := h.svc.Create(req)
	if errors.Is(err, service.ErrInvalidInput) {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, book)
}

// UpdateBook parses a JSON body and updates an existing book (PUT /api/v1/books/{id}).
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req model.UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	book, err := h.svc.Update(id, req)
	switch {
	case errors.Is(err, service.ErrBookNotFound):
		response.Error(w, http.StatusNotFound, "book not found")
	case errors.Is(err, service.ErrInvalidInput):
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
	default:
		response.JSON(w, http.StatusOK, book)
	}
}

// DeleteBook removes a book by ID (DELETE /api/v1/books/{id}).
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.svc.Delete(id); errors.Is(err, service.ErrBookNotFound) {
		response.Error(w, http.StatusNotFound, "book not found")
		return
	}

	// 204 No Content is the standard response for a successful delete.
	w.WriteHeader(http.StatusNoContent)
}
