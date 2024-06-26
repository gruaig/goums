// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package users

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name, lastname, email, username, password,  enabled
    ) VALUES (
        $1, $2, $3, $4, $5, $6
    )
RETURNING id, name, lastname, email, username, password, "CreatedAt", "UpdatedAt", enabled
`

type CreateUserParams struct {
	Name     sql.NullString
	Lastname sql.NullString
	Email    sql.NullString
	Username sql.NullString
	Password sql.NullString
	Enabled  sql.NullBool
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Lastname,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.Enabled,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Enabled,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, lastname, email, username, password, "CreatedAt", "UpdatedAt", enabled from users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Lastname,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Enabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getuser = `-- name: Getuser :one
SELECT id, name, lastname, email, username, password, "CreatedAt", "UpdatedAt", enabled FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) Getuser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getuser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Enabled,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, lastname, email, username, password, "CreatedAt", "UpdatedAt", enabled FROM users
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Lastname,
			&i.Email,
			&i.Username,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Enabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
