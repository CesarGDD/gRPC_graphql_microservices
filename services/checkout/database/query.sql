-- name: CreateOrder :one
INSERT INTO orders (user_id, total_price, status)
VALUES ($1, $2, 'Pending') RETURNING order_id;

-- name: AddOrderItem :exec
INSERT INTO order_items (order_id, product_id, quantity)
VALUES ($1, $2, $3);

-- name: GetOrderDetails :one
SELECT order_id, user_id, total_price, status
FROM orders
WHERE order_id = $1;

-- name: GetOrdersDetailsByUserId :many
SELECT order_id, user_id, total_price, status
FROM orders
WHERE user_id = $1;

-- name: GetOrderItems :many
SELECT product_id, quantity
FROM order_items
WHERE order_id = $1;

-- name: RecordPaymentTransaction :exec
INSERT INTO payment (order_id, payment_method)
VALUES ($1, $2);

-- name: GetPaymentDetails :one
SELECT payment_method
FROM payment
WHERE order_id = $1;