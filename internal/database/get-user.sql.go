// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: get-user.sql

package database

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, name FROM users
WHERE name = $1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
