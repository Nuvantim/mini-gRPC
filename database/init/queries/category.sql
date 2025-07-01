-- name: CreateCategory :one
INSERT INTO category (name) VALUES ($1) RETURNING *;
-- name: GetCategory :one
SELECT * FROM category WHERE id = $1 LIMIT 1;
-- name: ListCategory :many
SELECT * FROM category;
-- name: UpdateCategory :one
UPDATE category SET name = $2 WHERE id = $1 RETURNING *;
-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1; 