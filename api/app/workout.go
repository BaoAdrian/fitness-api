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
		log.Fatal("ERROR: Failed to read request body")
	}
	var workout db.Workout
	if err := json.Unmarshal(body, &workout); err != nil {
		log.Fatal("ERROR: Failed to unpack user data. Verify JSON formatting.")
	}

	db.AddWorkout(workout, app.Database)
}

// (GET) Endpoint: /workouts/exerciseid/{exerciseid}
func (app *App) getWorkoutsByExerciseID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["exerciseid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	results, err := db.GetWorkoutsByExerciseID(id, app.Database)
	if err = json.NewEncoder(w).Encode(results); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /workouts/exerciseid/{exerciseid}
func (app *App) deleteWorkoutsByExerciseID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["exerciseid"]

	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	db.DeleteWorkoutByExerciseID(id, app.Database)
}

// (GET) Endpoint: /workouts/routineid/{routineid}
func (app *App) getWorkoutsByRoutineID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["routineid"]
	if !ok {
		log.Fatal("ERROR: No id was provided")
	}

	workout, err := db.GetWorkoutsByRoutineID(id, app.Database)
	if err = json.NewEncoder(w).Encode(workout); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /workouts/routineid/{routineid}
func (app *App) deleteWorkoutsByRoutineID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["routineid"]
	if !ok {
		log.Fatal("ERROR: No id was provided")
	}

	db.DeleteWorkoutByRoutineID(id, app.Database)
}

// (GET) Endpoint: /workouts/ids?optional_params
func (app *App) getWorkoutByPKIDs(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	routineid := vars.Get("routineid")
	exerciseid := vars.Get("exerciseid")

	var err error
	var workouts db.Workouts
	var workout db.Workout
	if routineid != "" && exerciseid != "" {
		workout, err = db.GetWorkoutByPKIDs(routineid, exerciseid, app.Database)
		if err = json.NewEncoder(w).Encode(workout); err != nil {
			log.Fatal("ERROR: Failed to Encode JSON")
		}
	} else if routineid != "" {
		workouts, err = db.GetWorkoutsByRoutineID(routineid, app.Database)
		if err = json.NewEncoder(w).Encode(workouts); err != nil {
			log.Fatal("ERROR: Failed to Encode JSON")
		}
	} else if exerciseid != "" {
		workouts, err = db.GetWorkoutsByExerciseID(exerciseid, app.Database)
		if err = json.NewEncoder(w).Encode(workouts); err != nil {
			log.Fatal("ERROR: Failed to Encode JSON")
		}
	} else {
		log.Info("No IDs have been provided")
	}
}

// (DELETE) Endpoint: /workouts/ids?required_params
// Only deletes based on (routineid, exerciseid) PK, not either/or (enforced)
func (app *App) deleteWorkoutByPKIDs(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	routineid := vars.Get("routineid")
	exerciseid := vars.Get("exerciseid")

	if routineid != "" && exerciseid != "" {
		db.DeleteWorkoutByPKIDs(routineid, exerciseid, app.Database)
	} else {
		log.Info("Invalid number of IDs have been provided")
	}
}
