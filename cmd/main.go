package main

import (
	"database/sql"
	"log"

	"github.com/trademate-be/api/server"
	"github.com/trademate-be/internal/repository/sqlc/store"
	_ "github.com/trademate-be/pkg/background"
	"github.com/trademate-be/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := store.NewStore(conn)
	server := server.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
