-- name: CreateProduct :one
INSERT INTO product (name,description,category_id,price) VALUES ($1,$2,$3,$4) RETURNING *;
-- name: GetProduct :one
SELECT FROM product WHERE id = $1;
-- name: ListProduct :many
SELECT * FROM product;
-- name: UpdateProduct :one
UPDATE product SET name=$2, description=$3, category_id=$4, price=$5 WHERE id = $1 RETURNING *;
-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;

-- SELECT p.id, p.name, p.description, p.category_id, p.price, p.created_at
-- FROM "public".product p 
-- 	INNER JOIN "public".category c1 ON ( c1.id = p.category_id  )  