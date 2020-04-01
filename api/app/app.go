package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// App struct
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

// DefaultResponse Struct
type DefaultResponse struct {
	Message string `json:"message"`
}

// Category Struct
type Category struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// SetupRouter Creates Router & Maps Handler Functions for API
func (app *App) SetupRouter() {

	// Exercise endpoints
	app.Router.Methods("GET").Path("/exercises").HandlerFunc(app.getExercises)
	app.Router.Methods("POST").Path("/exercises").HandlerFunc(app.addExercise)
	app.Router.Methods("GET").Path("/exercises/names").HandlerFunc(app.getExerciseNames)
	app.Router.Methods("GET").Path("/exercises/name/{name}").HandlerFunc(app.getExerciseByName)
	app.Router.Methods("DELETE").Path("/exercises/name/{name}").HandlerFunc(app.deleteExerciseByName)
	app.Router.Methods("GET").Path("/exercises/categories").HandlerFunc(app.getExerciseCategories)
	app.Router.Methods("GET").Path("/exercises/id/{exerciseid}").HandlerFunc(app.getExerciseByID)
	app.Router.Methods("DELETE").Path("/exercises/id/{exerciseid}").HandlerFunc(app.deleteExerciseByID)
	app.Router.Methods("GET").Path("/exercises/category/{category}").HandlerFunc(app.getExerciseByCategory)

	app.Router.Methods("GET").Path("/workouts").HandlerFunc(app.getWorkouts)
	app.Router.Methods("POST").Path("/workouts").HandlerFunc(app.addWorkout)

	// Default Endpoint
	app.Router.HandleFunc("", notFound)
}

// DEFAULT
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
