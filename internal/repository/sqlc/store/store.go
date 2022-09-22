package store

import (
	"database/sql"

	repository "github.com/trademate-be/internal/repository/sqlc"
)

type Store interface {
	repository.Querier
}

type SQLStore struct {
	*repository.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: repository.New(db),
	}
}
