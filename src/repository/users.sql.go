// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (email, username, password, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING users."GUID", users.email, users.username, users.password, users.created_at, users.updated_at
`

type CreateUserParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.GUID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getByUsername = `-- name: GetByUsername :one
SELECT "email", "username", "password"
FROM users
WHERE username = $1
`

type GetByUsernameRow struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) GetByUsername(ctx context.Context, username string) (GetByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getByUsername, username)
	var i GetByUsernameRow
	err := row.Scan(&i.Email, &i.Username, &i.Password)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "users"
SET email = $1, username = $2, password = $3, updated_at = CURRENT_TIMESTAMP
WHERE username = $2
RETURNING "GUID"
`

type UpdateUserParams struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Email, arg.Username, arg.Password)
	var GUID int32
	err := row.Scan(&GUID)
	return GUID, err
}
