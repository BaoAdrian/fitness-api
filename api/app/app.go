package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

// Exercises struct
type Exercises struct {
	Exercises []Exercise `json:"exercises"`
}

// Names struct
type Names struct {
	Names []string `json:"names"`
}

// Categories Struct
type Categories struct {
	Categories []string `json:"categories"`
}

// SetupRouter Creates Router & Maps Handler Functions for API
func (app *App) SetupRouter() {

	api := app.Router.PathPrefix("/api/v1").Subrouter()

	api.Methods("GET").Path("/exercises").HandlerFunc(app.getExercises)
	api.Methods("GET").Path("/exercises/names").HandlerFunc(app.getExerciseNames)
	api.Methods("GET").Path("/exercises/categories").HandlerFunc(app.getExerciseCategories)
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

	collection := Exercises{}
	for rows.Next() {
		exercise := Exercise{}
		if err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description); err != nil {
			panic(err)
		}
		collection.Exercises = append(collection.Exercises, exercise)
	}
	json.NewEncoder(w).Encode(collection)

	if err = rows.Err(); err != nil {
		panic(err)
	}
}

// Endpoint: /exercises/names
// Response: Collection of all exercise names within the database
func (app *App) getExerciseNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := app.Database.Query(`SELECT name FROM exercises`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	collection := Names{}
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			panic(err)
		}
		collection.Names = append(collection.Names, name)
	}
	json.NewEncoder(w).Encode(collection)

	if err = rows.Err(); err != nil {
		panic(err)
	}
}

// Endpoint: /exercises/categories
// Response: Collection of all exercise categories within the database
func (app *App) getExerciseCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := app.Database.Query(`SELECT category FROM exercises`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	collection := Categories{}
	for rows.Next() {
		var category string
		if err = rows.Scan(&category); err != nil {
			panic(err)
		}
		collection.Categories = append(collection.Categories, category)
	}
	json.NewEncoder(w).Encode(collection)

	if err = rows.Err(); err != nil {
		panic(err)
	}
}

// Endpoint: /exercises/id/{exerciseid}
// Response: Retrieves exercise(s) with given id
// Assumption: No two exercises have the same id, therefore, returned JSON
// should have a single object
func (app *App) getExerciseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["exerciseid"]
	if !ok {
		fmt.Println("Not ok")
	}

	rows, err := app.Database.Query(fmt.Sprintf("SELECT * FROM exercises WHERE exerciseid = %s", id))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	exercise := Exercise{}
	for rows.Next() {
		if err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description); err != nil {
			panic(err)
		}
	}
	json.NewEncoder(w).Encode(exercise)

	if err = rows.Err(); err != nil {
		panic(err)
	}

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
