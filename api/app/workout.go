package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/BaoAdrian/fitness-api/api/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// (GET) Endpoint: /workouts
// Retrieves all workouts from the database
func (app *App) getWorkouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	workouts, err := db.GetWorkouts(app.Database)
	if err = json.NewEncoder(w).Encode(workouts); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (POST) Endpoint: /workouts
func (app *App) addWorkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	var workout db.Workout
	if err := json.Unmarshal(body, &workout); err != nil {
		panic(err)
	}

	db.AddWorkout(workout, app.Database)
}

// (GET) Endpoint: /workouts/id/{workoutid}
func (app *App) getWorkoutByWorkoutID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["workoutid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	results, err := db.GetWorkoutByWorkoutID(id, app.Database)
	if err = json.NewEncoder(w).Encode(results); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /workouts/id/{workoutid}
func (app *App) deleteWorkoutByWorkoutID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["workoutid"]

	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	db.DeleteWorkoutByID(id, app.Database)
}

// (GET) Endpoint: /workouts/name/{name}
func (app *App) getWorkoutByWorkoutName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		log.Fatal("ERROR: No name was provided")
	}

	workout, err := db.GetWorkoutByName(name, app.Database)
	if err = json.NewEncoder(w).Encode(workout); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /workouts/name/{name}
func (app *App) deleteWorkoutByWorkoutName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		log.Fatal("ERROR: No name was provided")
	}

	db.DeleteWorkoutByName(name, app.Database)
}
