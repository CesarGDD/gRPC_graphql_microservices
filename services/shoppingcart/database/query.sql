-- name: AddItemToCart :exec
INSERT INTO cart_items (user_id, product_id, quantity)
VALUES ($1, $2, $3);

-- name: UpdateCartItem :exec
UPDATE cart_items SET quantity = $2
WHERE user_id = $1 AND product_id = $3;

-- name: RemoveItemFromCart :exec
DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2;

-- name: GetCartItems :many
SELECT product_id, quantity
FROM cart_items
WHERE user_id = $1;

-- name: ClearCart :exec
DELETE FROM cart_items WHERE user_id = $1;