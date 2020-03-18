package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App struct
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

// Exercise struct
type Exercise struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Category    string         `json:"category"`
	Description sql.NullString `json:"description"`
}

// Collection struct
type Collection struct {
	Collection []Exercise `json:"collection"`
}

// SetupRouter Creates Router & Maps Handler Functions for API
func (app *App) SetupRouter() {

	api := app.Router.PathPrefix("/api/v1").Subrouter()

	api.Methods("GET").Path("/exercises").HandlerFunc(app.getExercises)
	api.Methods("GET").Path("/exercises/names").HandlerFunc(app.getExerciseNames)
	api.Methods("GET").Path("/exercises/category").HandlerFunc(app.getExerciseCategories)
	api.Methods("GET").Path("/exercises/id/{exerciseid}").HandlerFunc(app.getExerciseByID)
	api.Methods("GET").Path("/exercises/name/{name}").HandlerFunc(app.getExerciseByName)
	api.Methods("GET").Path("/exercises/category/{category}").HandlerFunc(app.getExerciseByCategory)
	api.HandleFunc("", notFound)

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// Endpoint: /exercises
// Response: Collection of all exercises within the database
func (app *App) getExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := app.Database.Query(`SELECT * FROM exercises`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	collection := Collection{}
	for rows.Next() {
		exercise := Exercise{}
		if err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description); err != nil {
			panic(err)
		}
		collection.Collection = append(collection.Collection, exercise)
	}
	json.NewEncoder(w).Encode(collection)

	if err = rows.Err(); err != nil {
		panic(err)
	}
}

func (app *App) getExerciseNames(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (app *App) getExerciseCategories(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (app *App) getExerciseByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (app *App) getExerciseByName(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (app *App) getExerciseByCategory(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// DEFAULT
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
