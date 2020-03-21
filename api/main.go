package main

import (
	"github.com/BaoAdrian/fitness-api/api/app"
	"github.com/BaoAdrian/fitness-api/api/db"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		panic(err)
	}

	app := &app.App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	app.SetupRouter()
}
