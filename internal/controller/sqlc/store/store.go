package store

import (
	"database/sql"

	controller "github.com/trademate-be/internal/controller/sqlc"
)

type Store interface {
	controller.Querier
}

type SQLStore struct {
	*controller.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: controller.New(db),
	}
}
