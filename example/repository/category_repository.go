package repository

import (
	"context"
	"database/sql"
	"errors"

	"example/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, name string) (*models.Category, error) {
	const query = `
		INSERT INTO category (name)
		VALUES ($1, $2)
		RETURNING id, name
	`

	var category models.Category
	err := r.db.QueryRowxContext(ctx, query, name).StructScan(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id int64) (*models.Category, error) {
	const query = `
		SELECT id, name
		FROM category
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
		SELECT id, name
		FROM category
		ORDER BY created_at DESC
	`

	var categories []models.Category
	err := r.db.SelectContext(ctx, &categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id int64, name string) (*models.Category, error) {
	const query = `
		UPDATE category
		SET name = $2 updated_at = NOW()
		WHERE id = $1
		RETURNING id, name
	`

	var category models.Category
	err := r.db.QueryRowxContext(ctx, query, id, name).StructScan(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id int64) (bool, error) {
	const query = `
		DELETE FROM category
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
