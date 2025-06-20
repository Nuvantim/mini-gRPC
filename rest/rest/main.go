package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"micro/internal/database"
	"micro/internal/domain/repository"
)

type CategoryHandler struct {
	queries *repository.Queries
}

func main() {
	database.InitDB()
	defer database.CloseDB()

	handler := &CategoryHandler{
		queries: database.Queries,
	}

	// Initialize the router
	r := mux.NewRouter()

	// Setup routes using mux router
	r.HandleFunc("/categories", handler.handleCategories).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/categories/{id:[0-9]+}", handler.handleCategory).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	// Start server
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// CategoryRequest represents the JSON structure for create/update requests
type CategoryRequest struct {
	Name string `json:"name"`
}

// CategoryResponse represents the JSON structure for responses
type CategoryResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// handleCategories handles GET (list) and POST (create) operations for /categories
func (h *CategoryHandler) handleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.ListCategory(w, r)
	case http.MethodPost:
		h.createCategory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleCategory handles GET, PUT, and DELETE operations for /categories/{id}
func (h *CategoryHandler) handleCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getCategory(w, r, int32(id))
	case http.MethodPut:
		h.updateCategory(w, r, int32(id))
	case http.MethodDelete:
		h.deleteCategory(w, r, int32(id))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ListCategory returns all categories
func (h *CategoryHandler) ListCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := h.queries.ListCategory(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		return
	}

	response := make([]CategoryResponse, len(categories))
	for i, cat := range categories {
		response[i] = CategoryResponse{
			ID:        cat.ID,
			Name:      cat.Name,
			CreatedAt: cat.CreatedAt,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getCategory returns a single category by ID
func (h *CategoryHandler) getCategory(w http.ResponseWriter, r *http.Request, id int32) {
	category, err := h.queries.GetCategory(context.Background(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch category", http.StatusInternalServerError)
		}
		return
	}

	response := CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// createCategory creates a new category
func (h *CategoryHandler) createCategory(w http.ResponseWriter, r *http.Request) {
	var req CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(req)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Start transaction
	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(context.Background())

	qtx := h.queries.WithTx(tx)

	category, err := qtx.CreateCategory(context.Background(), req.Name)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	response := CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// updateCategory updates an existing category
func (h *CategoryHandler) updateCategory(w http.ResponseWriter, r *http.Request, id int32) {
	var req CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Start transaction
	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(context.Background())

	qtx := h.queries.WithTx(tx)

	category, err := qtx.UpdateCategory(context.Background(), repository.UpdateCategoryParams{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update category", http.StatusInternalServerError)
		}
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	response := CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// deleteCategory deletes a category
func (h *CategoryHandler) deleteCategory(w http.ResponseWriter, r *http.Request, id int32) {
	// Start transaction
	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback(context.Background())

	qtx := h.queries.WithTx(tx)

	err = qtx.DeleteCategory(context.Background(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Category not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		}
		return
	}

	if err := tx.Commit(context.Background()); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
