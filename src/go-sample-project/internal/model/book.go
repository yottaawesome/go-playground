// Package model defines the core data types used across the application.
// These structs are intentionally kept free of framework-specific annotations
// to stay decoupled from any particular HTTP library or database driver.
package model

import "time"

// Book represents a single book in the system.
type Book struct {
	ID        string    `json:"id"`         // Unique identifier (UUID).
	Title     string    `json:"title"`      // Title of the book.
	Author    string    `json:"author"`     // Author's full name.
	ISBN      string    `json:"isbn"`       // International Standard Book Number.
	Pages     int       `json:"pages"`      // Total page count.
	CreatedAt time.Time `json:"created_at"` // Timestamp of creation.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of last update.
}

// CreateBookRequest carries the fields required to create a new book.
// Validation is performed at the handler/service boundary.
type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Pages  int    `json:"pages"`
}

// UpdateBookRequest carries the fields that can be modified on an existing book.
type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Pages  int    `json:"pages"`
}
