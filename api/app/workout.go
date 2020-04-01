package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/BaoAdrian/fitness-api/api/db"
)

// (GET) Endpoint: /workouts
// Retrieves all workouts from the database
func (app *App) getWorkouts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	workouts, err := db.GetWorkouts(app.Database)
	if err != nil {
		if err := json.NewEncoder(w).Encode(DefaultResponse{Message: "No data found."}); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(workouts); err != nil {
			panic(err)
		}
	}
}

// (POST) Endpoint: /workouts
// Posts a workout with provided JSON
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
