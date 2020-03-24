package main

import (
	"log"
	"net/http"

	"github.com/BaoAdrian/fitness-api/api/app"
	"github.com/BaoAdrian/fitness-api/api/db"

	"github.com/gorilla/mux"
)

func main() {
	// Create instance from database container (container_name: 'db')
	database, err := db.CreateDatabase("db:3306")
	if err != nil {
		panic(err)
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
