// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: journals.sql

package db

import (
	"context"
)

const createJournal = `-- name: CreateJournal :one
INSERT INTO journals (
    user_id,
    title, 
    strategy, 
    initial_balance,
    current_balance, 
    risk_tolerance 
) VALUES (
    $1, $2, $3, $4, $5, $6 
) 
RETURNING id, user_id, title, strategy, initial_balance, current_balance, risk_tolerance, created_at
`

type CreateJournalParams struct {
	UserID         int64
	Title          string
	Strategy       string
	InitialBalance string
	CurrentBalance string
	RiskTolerance  string
}

func (q *Queries) CreateJournal(ctx context.Context, arg CreateJournalParams) (Journal, error) {
	row := q.db.QueryRowContext(ctx, createJournal,
		arg.UserID,
		arg.Title,
		arg.Strategy,
		arg.InitialBalance,
		arg.CurrentBalance,
		arg.RiskTolerance,
	)
	var i Journal
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Strategy,
		&i.InitialBalance,
		&i.CurrentBalance,
		&i.RiskTolerance,
		&i.CreatedAt,
	)
	return i, err
}

const deleteJournal = `-- name: DeleteJournal :exec
DELETE FROM journals WHERE id = $1
`

func (q *Queries) DeleteJournal(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteJournal, id)
	return err
}

const getJournal = `-- name: GetJournal :one
SELECT id, user_id, title, strategy, initial_balance, current_balance, risk_tolerance, created_at FROM journals
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetJournal(ctx context.Context, id int64) (Journal, error) {
	row := q.db.QueryRowContext(ctx, getJournal, id)
	var i Journal
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Strategy,
		&i.InitialBalance,
		&i.CurrentBalance,
		&i.RiskTolerance,
		&i.CreatedAt,
	)
	return i, err
}

const listJournals = `-- name: ListJournals :many
SELECT id, user_id, title, strategy, initial_balance, current_balance, risk_tolerance, created_at FROM journals 
ORDER BY id 
LIMIT $1 
OFFSET $2
`

type ListJournalsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListJournals(ctx context.Context, arg ListJournalsParams) ([]Journal, error) {
	rows, err := q.db.QueryContext(ctx, listJournals, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Journal{}
	for rows.Next() {
		var i Journal
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Strategy,
			&i.InitialBalance,
			&i.CurrentBalance,
			&i.RiskTolerance,
			&i.CreatedAt,
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

const updateJournal = `-- name: UpdateJournal :one
UPDATE journals 
SET user_id = $2, 
    title = $3, 
    strategy = $4, 
    initial_balance = $5, 
    current_balance = $6, 
    risk_tolerance = $7 
WHERE id = $1
RETURNING id, user_id, title, strategy, initial_balance, current_balance, risk_tolerance, created_at
`

type UpdateJournalParams struct {
	ID             int64
	UserID         int64
	Title          string
	Strategy       string
	InitialBalance string
	CurrentBalance string
	RiskTolerance  string
}

func (q *Queries) UpdateJournal(ctx context.Context, arg UpdateJournalParams) (Journal, error) {
	row := q.db.QueryRowContext(ctx, updateJournal,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.Strategy,
		arg.InitialBalance,
		arg.CurrentBalance,
		arg.RiskTolerance,
	)
	var i Journal
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Strategy,
		&i.InitialBalance,
		&i.CurrentBalance,
		&i.RiskTolerance,
		&i.CreatedAt,
	)
	return i, err
}
