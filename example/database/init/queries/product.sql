-- name: CreateProduct :one
INSERT INTO product (name,description,category_id,price) VALUES ($1,$2,$3,$4) RETURNING *;

-- name: GetProduct :one
SELECT sqlc.embed(product), sqlc.embed(category)
FROM "public".product
INNER JOIN "public".category ON ( category.id = product.category_id  ) WHERE product.id = $1 LIMIT 1;

-- name: ListProduct :many
SELECT * FROM product;

-- name: UpdateProduct :one
UPDATE product SET name=$2, description=$3, category_id=$4, price=$5 WHERE id = $1 RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;


-- SELECT * FROM product WHERE id = $1 LIMIT 1;
