-- name: CreateCategory :one
INSERT INTO category (id,name) VALUES ($1,$2) RETURNING *;
-- name: GetCategory :one
SELECT * FROM category WHERE id = $1;
-- name: ListCategory :many
SELECT * FROM category ORDER BY name;
-- name: UpdateCategory :one
UPDATE category SET name= $2 WHERE id = $1 RETURNING *;
-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1; 