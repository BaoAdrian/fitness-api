package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Endpoint: /exercises
// Response: Collection of all exercises within the database
func exercises(w http.ResponseWriter, r *http.Request) {
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

	db, err := sql.Open("mysql", "root:password@tcp(db:3306)/fitnessdb")
	if err != nil {
		fmt.Println("[ERROR]: Unable to connect to database!")
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM exercises`)
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
	defer db.Close()
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

func main() {
	r := mux.NewRouter()

	// Using subrouter for easier migration from API v1 to v2
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/exercises", exercises).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}
