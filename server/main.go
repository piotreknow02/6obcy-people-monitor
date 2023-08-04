package main

import (
	"fmt"
	"os"

	"github.com/piotreknow02/6obcy-people-monitor/server/config"
	"github.com/piotreknow02/6obcy-people-monitor/server/database"
	"github.com/piotreknow02/6obcy-people-monitor/server/router"
)

func main() {
	config.Load()

	db, err := database.Setup()
	if err != nil {
		fmt.Println(fmt.Errorf("Error setting up database: %s", err.Error()))
		os.Exit(2)
	}
	r := router.Setup(db)

	r.Listen(fmt.Sprintf(":%d", config.Conf.Port))
}
