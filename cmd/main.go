package main

import (
	"database/sql"
	"log"

	server "github.com/trademate-be/api"
	db "github.com/trademate-be/internal/db/sqlc"
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

	store := db.NewStore(conn)
	server := server.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
