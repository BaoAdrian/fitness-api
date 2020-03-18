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

// SetupRouter Creates Router & Maps Handler Functions for API
func (app *App) SetupRouter() {

	api := app.Router.PathPrefix("/api/v1").Subrouter()

	api.Methods("GET").Path("/exercises").HandlerFunc(app.exercises)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// Endpoint: /exercises
// Response: Collection of all exercises within the database
func (app *App) exercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	type Exercise struct {
		ID          int            `json:"id"`
		Name        string         `json:"name"`
		Category    string         `json:"category"`
		Description sql.NullString `json:"description"`
	}

	type Collection struct {
		Collection []Exercise `json:"collection"`
	}

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

// POST
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "POST called"}`))
}

// PUT
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "PUT called"}`))
}

// DELETE
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "DELETE called"}`))
}

// DEFAULT
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
