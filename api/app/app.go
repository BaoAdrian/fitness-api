package app

import (
	"database/sql"
	"net/http"

	"github.com/BaoAdrian/fitness-api/api/db"
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

func setupApp() (app *App) {
	database, err := db.CreateDatabase("localhost:3306")
	if err != nil {
		panic(err)
	}

	app = &App{
		Router:   mux.NewRouter(),
		Database: database,
	}
	return app
}

// SetupRouter Creates Router & Maps Handler Functions for API
func (app *App) SetupRouter() {

	// Routines Endpoints
	app.Router.Methods("GET").Path("/routines").HandlerFunc(app.getRoutines)
	app.Router.Methods("POST").Path("/routines").HandlerFunc(app.addRoutine)
	app.Router.Methods("GET").Path("/routines/routineid/{routineid}").HandlerFunc(app.getRoutineByRoutineID)
	app.Router.Methods("DELETE").Path("/routines/routineid/{routineid}").HandlerFunc(app.deleteRoutineByRoutineID)
	app.Router.Methods("GET").Path("/routines/userid/{userid}").HandlerFunc(app.getRoutinesByUserID)
	app.Router.Methods("DELETE").Path("/routines/userid/{userid}").HandlerFunc(app.deleteRoutinesByUserID)

	// Users Endpoints
	app.Router.Methods("GET").Path("/users").HandlerFunc(app.getUsers)
	app.Router.Methods("POST").Path("/users").HandlerFunc(app.addUser)
	app.Router.Methods("GET").Path("/users/id/{userid}").HandlerFunc(app.getUserByUserID)
	app.Router.Methods("DELETE").Path("/users/id/{userid}").HandlerFunc(app.deleteUserByUserID)

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

	// Workouts Endpoints
	app.Router.Methods("GET").Path("/workouts").HandlerFunc(app.getWorkouts)
	app.Router.Methods("POST").Path("/workouts").HandlerFunc(app.addWorkout)
	app.Router.Methods("GET").Path("/workouts/exerciseid/{exerciseid}").HandlerFunc(app.getWorkoutsByExerciseID)
	app.Router.Methods("DELETE").Path("/workouts/exerciseid/{exerciseid}").HandlerFunc(app.deleteWorkoutsByExerciseID)
	app.Router.Methods("GET").Path("/workouts/routineid/{routineid}").HandlerFunc(app.getWorkoutsByRoutineID)
	app.Router.Methods("DELETE").Path("/workouts/routineid/{routineid}").HandlerFunc(app.deleteWorkoutsByRoutineID)
	app.Router.Methods("GET").Path("/workouts/ids").HandlerFunc(app.getWorkoutByPKIDs)
	app.Router.Methods("DELETE").Path("/workouts/ids").HandlerFunc(app.deleteWorkoutByPKIDs)

	// Default Endpoint
	app.Router.HandleFunc("", notFound)
}

// DEFAULT
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
