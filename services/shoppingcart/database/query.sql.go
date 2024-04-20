// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
)

const addItemToCart = `-- name: AddItemToCart :exec
INSERT INTO cart_items (user_id, product_id, quantity)
VALUES ($1, $2, $3)
`

type AddItemToCartParams struct {
	UserID    int32
	ProductID int32
	Quantity  int32
}

func (q *Queries) AddItemToCart(ctx context.Context, arg AddItemToCartParams) error {
	_, err := q.db.Exec(ctx, addItemToCart, arg.UserID, arg.ProductID, arg.Quantity)
	return err
}

const clearCart = `-- name: ClearCart :exec
DELETE FROM cart_items WHERE user_id = $1
`

func (q *Queries) ClearCart(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, clearCart, userID)
	return err
}

const getCartItems = `-- name: GetCartItems :many
SELECT product_id, quantity
FROM cart_items
WHERE user_id = $1
`

type GetCartItemsRow struct {
	ProductID int32
	Quantity  int32
}

func (q *Queries) GetCartItems(ctx context.Context, userID int32) ([]GetCartItemsRow, error) {
	rows, err := q.db.Query(ctx, getCartItems, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCartItemsRow
	for rows.Next() {
		var i GetCartItemsRow
		if err := rows.Scan(&i.ProductID, &i.Quantity); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeItemFromCart = `-- name: RemoveItemFromCart :exec
DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2
`

type RemoveItemFromCartParams struct {
	UserID    int32
	ProductID int32
}

func (q *Queries) RemoveItemFromCart(ctx context.Context, arg RemoveItemFromCartParams) error {
	_, err := q.db.Exec(ctx, removeItemFromCart, arg.UserID, arg.ProductID)
	return err
}

const updateCartItem = `-- name: UpdateCartItem :exec
UPDATE cart_items SET quantity = $2
WHERE user_id = $1 AND product_id = $3
`

type UpdateCartItemParams struct {
	UserID    int32
	Quantity  int32
	ProductID int32
}

func (q *Queries) UpdateCartItem(ctx context.Context, arg UpdateCartItemParams) error {
	_, err := q.db.Exec(ctx, updateCartItem, arg.UserID, arg.Quantity, arg.ProductID)
	return err
}
