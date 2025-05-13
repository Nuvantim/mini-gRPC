-- name: GetCategory :many
SELECT * FROM categories;
-- name: FindCategory :one
SELECT * FROM categories WHERE id = $1;
-- name: CreateCatgory :exec
INSERT INTO categories (name) VALUES ($1) RETURNING *;
-- name: UpdateCategory :exec
UPDATE categories SET name = $1
WHERE id = $2 RETURNING *;
-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;