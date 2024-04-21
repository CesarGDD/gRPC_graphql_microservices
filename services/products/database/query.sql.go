// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products (user_id, name, url, title, description, price)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateProductParams struct {
	UserID      int32
	Name        string
	Url         string
	Title       string
	Description string
	Price       int32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.Exec(ctx, createProduct,
		arg.UserID,
		arg.Name,
		arg.Url,
		arg.Title,
		arg.Description,
		arg.Price,
	)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID int32) error {
	_, err := q.db.Exec(ctx, deleteProduct, productID)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, user_id, name, url, title, description, price FROM products WHERE product_id = $1
`

func (q *Queries) GetProduct(ctx context.Context, productID int32) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, productID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const getProductByName = `-- name: GetProductByName :one
SELECT product_id, user_id, name, url, title, description, price FROM products WHERE name = $1
`

func (q *Queries) GetProductByName(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRow(ctx, getProductByName, name)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const getProductByUserId = `-- name: GetProductByUserId :many
SELECT product_id, user_id, name, url, title, description, price FROM products WHERE user_id = $1
`

func (q *Queries) GetProductByUserId(ctx context.Context, userID int32) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProductByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Title,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProducts = `-- name: GetProducts :many
SELECT product_id, user_id, name, url, title, description, price FROM products
`

// Get All Products
func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.UserID,
			&i.Name,
			&i.Url,
			&i.Title,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products SET name = $2, url = $3, description = $4, price = $5
WHERE product_id = $1
`

type UpdateProductParams struct {
	ProductID   int32
	Name        string
	Url         string
	Description string
	Price       int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.ProductID,
		arg.Name,
		arg.Url,
		arg.Description,
		arg.Price,
	)
	return err
}
