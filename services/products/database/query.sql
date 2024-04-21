-- name: CreateProduct :exec
INSERT INTO products (user_id, name, url, title, description, price)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetProductByName :one
SELECT * FROM products WHERE name = $1;

-- name: GetProductByUserId :many
SELECT * FROM products WHERE user_id = $1;

-- name: UpdateProduct :exec
UPDATE products SET name = $2, url = $3, description = $4, price = $5
WHERE product_id = $1;

-- name: GetProduct :one
SELECT * FROM products WHERE product_id = $1;

-- Get All Products
-- name: GetProducts :many
SELECT * FROM products;

-- name: DeleteProduct :exec
DELETE FROM products WHERE product_id = $1;
