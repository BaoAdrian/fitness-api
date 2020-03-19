package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

	rows, err := app.RunQuery(`SELECT * FROM exercises`)

	collection := Exercises{}
	for rows.Next() {
		exercise := Exercise{}
		if err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Exercises = append(collection.Exercises, exercise)
	}

	// Write output
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

// Endpoint: /exercises/names
// Response: Collection of all exercise names within the database
func (app *App) getExerciseNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := app.RunQuery(`SELECT name FROM exercises`)

	collection := Names{}
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Names = append(collection.Names, name)
	}

	// Write output
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
	}
}

// Endpoint: /exercises/categories
// Response: Collection of all exercise categories within the database
func (app *App) getExerciseCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := app.RunQuery(`SELECT category FROM exercises`)

	collection := Categories{}
	for rows.Next() {
		var category string
		if err = rows.Scan(&category); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Categories = append(collection.Categories, category)
	}

	// Write output
	if err := json.NewEncoder(w).Encode(collection); err != nil {
		panic(err)
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

	exercise := Exercise{}
	err := app.Database.QueryRow(fmt.Sprintf("SELECT * FROM exercises WHERE exerciseid = %s", id)).Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description)
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

	exercise := Exercise{}
	err := app.Database.QueryRow(fmt.Sprintf(`SELECT * FROM exercises WHERE name = "%s"`, name)).Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description)
	if err != nil {
		log.Warn("Database SELECT failed")
		json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."})
	} else {
		json.NewEncoder(w).Encode(exercise)
	}
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

// RunQuery Runs specified query on database
func (app *App) RunQuery(query string) (*sql.Rows, error) {
	rows, err := app.Database.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
