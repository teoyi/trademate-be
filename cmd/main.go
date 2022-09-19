package main

import (
	"fmt"
	"log"

	"github.com/trademate-be/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	fmt.Println("%v", config)
}
