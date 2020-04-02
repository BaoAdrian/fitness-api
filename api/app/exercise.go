package app

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/BaoAdrian/fitness-api/api/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// (GET) Endpoint: /exercises
func (app *App) getExercises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	exercises, err := db.GetExercises(app.Database)
	if err = json.NewEncoder(w).Encode(exercises); err != nil {
		panic(err)
	}
}

// (GET) Endpoint: /exercises/names
func (app *App) getExerciseNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	names, err := db.GetExerciseNames(app.Database)
	if err = json.NewEncoder(w).Encode(names); err != nil {
		panic(err)
	}
}

// (GET) Endpoint: /exercises/categories
// Response: Collection of all exercise categories within the database &
// their associated count
func (app *App) getExerciseCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	categories, err := db.GetExerciseCategories(app.Database)
	if err = json.NewEncoder(w).Encode(categories); err != nil {
		panic(err)
	}
}

// (GET) Endpoint: /exercises/id/{exerciseid}
// Response: Retrieves exercise with given 'exerciseid'
// Assertion: Since the database defines a constraint for unique ids, the
// query is guaranteed to retrieve <= 1 record.
func (app *App) getExerciseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["exerciseid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	exercise, err := db.GetExerciseByID(id, app.Database)
	if err = json.NewEncoder(w).Encode(exercise); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (GET) Endpoint: /exercises/name/{name}
// Response: Retrieves exercise with a given 'name'
// Assertion: Since the database defines a constraint for unique names, the
// query is guaranteed to retrieve <= 1 record.
func (app *App) getExerciseByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		log.Fatal("ERROR: No name was provided")
	}

	exercise, err := db.GetExerciseByName(name, app.Database)
	if err = json.NewEncoder(w).Encode(exercise); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (GET) Endpoint: /exercises/category/{category}
func (app *App) getExerciseByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	category, ok := vars["category"]
	if !ok {
		log.Fatal("ERROR: No category was provided")
	}

	collection, err := db.GetExerciseByCategory(category, app.Database)
	if err = json.NewEncoder(w).Encode(collection); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (POST) Endpoint: /exercises
func (app *App) addExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	var exercise db.Exercise
	if err := json.Unmarshal(body, &exercise); err != nil {
		panic(err)
	}

	db.AddExercise(exercise, app.Database)
}

// (DELETE) Endpoint: /exercises/id/{exerciseid}
func (app *App) deleteExerciseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["exerciseid"]

	if !ok {
		log.Fatal("ERROR: No id was provided")
	}

	db.DeleteExerciseByID(id, app.Database)
}

// (DELETE) Endpoint: /exercises/name/{name}
func (app *App) deleteExerciseByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	name, ok := vars["name"]

	if !ok {
		log.Fatal("ERROR: No name was provided")
	}

	db.DeleteExerciseByName(name, app.Database)
}

// (GET) Endpoint: /exercises/workoutid/{workoutid}
// Retrieves all collection of exercises listed under a specific workout
func (app *App) getExercisesByWorkoutID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	workoutid, ok := vars["workoutid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	exerciseIds := db.GetExerciseIDByWorkoutID(workoutid, app.Database)
	exercises := db.Exercises{}

	for _, exerciseid := range exerciseIds {
		exercise, err := db.GetExerciseByID(strconv.Itoa(exerciseid), app.Database)
		if err != nil {
			log.Fatal("ERROR: Failed to retrieve Exercise with id: " + string(exerciseid))
		}
		exercises.Exercises = append(exercises.Exercises, exercise)
	}

	json.NewEncoder(w).Encode(exercises)
}
