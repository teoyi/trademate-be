package main

import (
	"database/sql"
	"fmt"
	"log"

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

	fmt.Println(conn)
}
