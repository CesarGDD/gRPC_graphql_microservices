-- name: CreateUser :one
INSERT INTO users (username, password_hash, role)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetUserByUsername :one
SELECT user_id, username, password_hash, role FROM users WHERE username = $1;

-- name: UpdateUser :exec
UPDATE users SET username = $2, password_hash = $3, role = $4
WHERE user_id = $1;

-- name: GetUser :one
SELECT user_id, username, role FROM users WHERE user_id = $1;

-- Get All Users
-- name: GetUsers :many
SELECT user_id, username, role FROM users;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;

