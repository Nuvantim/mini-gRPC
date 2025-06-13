package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"example/models"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, name, description string) (*models.Category, error) {
	const query = `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description, created_at, updated_at
	`

	var category models.Category
	err := r.db.QueryRowxContext(ctx, query, name, description).StructScan(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id string) (*models.Category, error) {
	const query = `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		WHERE id = $1
	`

	var category models.Category
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) List(ctx context.Context) ([]models.Category, error) {
	const query = `
		SELECT id, name, description, created_at, updated_at
		FROM categories
		ORDER BY created_at DESC
	`

	var categories []models.Category
	err := r.db.SelectContext(ctx, &categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id, name, description string) (*models.Category, error) {
	const query = `
		UPDATE categories
		SET name = $2, description = $3, updated_at = NOW()
		WHERE id = $1
		RETURNING id, name, description, created_at, updated_at
	`

	var category models.Category
	err := r.db.QueryRowxContext(ctx, query, id, name, description).StructScan(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) (bool, error) {
	const query = `
		DELETE FROM categories
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
