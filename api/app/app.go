package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/BaoAdrian/fitness-api/api/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

	exercises, err := db.GetExercises(app.Database)
	if err != nil {
		if err := json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(exercises); err != nil {
			panic(err)
		}
	}
}

// Endpoint: /exercises/names
// Response: Collection of all exercise names within the database
func (app *App) getExerciseNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	names, err := db.GetExerciseNames(app.Database)
	if err != nil {
		if err := json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(names); err != nil {
			panic(err)
		}
	}
}

// Endpoint: /exercises/categories
// Response: Collection of all exercise categories within the database &
// their associated count
func (app *App) getExerciseCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	categories, err := db.GetExerciseCategories(app.Database)
	if err != nil {
		if err := json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			panic(err)
		}
	}
}

// Endpoint: /exercises/id/{exerciseid}
// Response: Retrieves exercise with given id
// Assertion: Since the database defines a constraint for unique ids, the
// query is guaranteed to retrieve <= 1 record.
func (app *App) getExerciseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["exerciseid"]
	if !ok {
		log.Fatal("No ID was provided")
	}

	exercise, err := db.GetExerciseByID(id, app.Database)
	if err != nil {
		log.Fatal("Database SELECT failed")
		json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."})
	} else {
		json.NewEncoder(w).Encode(exercise)
	}
}

// Endpoint: /exercises/name/{name}
// Response: Retrieves exercise with a given name
// Assertion: Since the database defines a constraint for unique names, the
// query is guaranteed to retrieve <= 1 record.
func (app *App) getExerciseByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		log.Fatal("No name was provided")
	}

	exercise, err := db.GetExerciseByName(name, app.Database)
	if err != nil {
		log.Warn("Database SELECT failed")
		json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."})
	} else {
		json.NewEncoder(w).Encode(exercise)
	}
}

// Endpoint: /exercises/category/{category}
// Response: Retrieves exercise associated with given category
func (app *App) getExerciseByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	category, ok := vars["category"]
	if !ok {
		log.Fatal("No name was provided")
	}

	collection, err := db.GetExerciseByCategory(category, app.Database)
	if err != nil {
		if err := json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(collection); err != nil {
			panic(err)
		}
	}
}

// DEFAULT
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
