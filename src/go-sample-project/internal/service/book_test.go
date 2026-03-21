// Package service tests verify the business logic in isolation—no HTTP
// server, no router, just pure Go function calls.
package service

import (
	"testing"

	"github.com/example/go-sample-project/internal/model"
)

// TestCreateAndGet verifies that a book can be created and then retrieved.
func TestCreateAndGet(t *testing.T) {
	svc := NewBookService()

	req := model.CreateBookRequest{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan & Brian Kernighan",
		ISBN:   "978-0134190440",
		Pages:  380,
	}

	created, err := svc.Create(req)
	if err != nil {
		t.Fatalf("unexpected error creating book: %v", err)
	}

	if created.ID == "" {
		t.Fatal("expected non-empty ID")
	}

	// Retrieve the same book by its ID.
	got, err := svc.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error getting book: %v", err)
	}

	if got.Title != req.Title {
		t.Errorf("title = %q, want %q", got.Title, req.Title)
	}
}

// TestCreateValidation ensures that missing required fields are rejected.
func TestCreateValidation(t *testing.T) {
	svc := NewBookService()

	// Missing both title and author.
	_, err := svc.Create(model.CreateBookRequest{})
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}

// TestDeleteRemovesBook confirms that a deleted book can no longer be found.
func TestDeleteRemovesBook(t *testing.T) {
	svc := NewBookService()

	book, _ := svc.Create(model.CreateBookRequest{
		Title:  "Concurrency in Go",
		Author: "Katherine Cox-Buday",
	})

	if err := svc.Delete(book.ID); err != nil {
		t.Fatalf("unexpected error deleting book: %v", err)
	}

	_, err := svc.GetByID(book.ID)
	if err != ErrBookNotFound {
		t.Errorf("expected ErrBookNotFound, got %v", err)
	}
}

// TestListReturnsAll creates several books and verifies List returns them all.
func TestListReturnsAll(t *testing.T) {
	svc := NewBookService()

	titles := []string{"Book A", "Book B", "Book C"}
	for _, title := range titles {
		_, _ = svc.Create(model.CreateBookRequest{Title: title, Author: "Author"})
	}

	books := svc.List()
	if len(books) != len(titles) {
		t.Errorf("got %d books, want %d", len(books), len(titles))
	}
}

// TestUpdateModifiesFields verifies that an update changes the book's fields.
func TestUpdateModifiesFields(t *testing.T) {
	svc := NewBookService()

	book, _ := svc.Create(model.CreateBookRequest{
		Title:  "Original Title",
		Author: "Original Author",
	})

	updated, err := svc.Update(book.ID, model.UpdateBookRequest{
		Title:  "New Title",
		Author: "New Author",
		ISBN:   "123-456",
		Pages:  200,
	})
	if err != nil {
		t.Fatalf("unexpected error updating book: %v", err)
	}

	if updated.Title != "New Title" {
		t.Errorf("title = %q, want %q", updated.Title, "New Title")
	}

	// CreatedAt should not change; UpdatedAt should be at least as recent.
	if updated.UpdatedAt.Before(updated.CreatedAt) {
		t.Error("expected UpdatedAt to be >= CreatedAt")
	}

	if updated.CreatedAt != book.CreatedAt {
		t.Error("expected CreatedAt to remain unchanged after update")
	}
}
