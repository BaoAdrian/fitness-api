package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/BaoAdrian/fitness-api/api/db"
	"github.com/stretchr/testify/assert"
)

var dummyWorkout = db.Workout{
	ExerciseID: 1,
	RoutineID:  12345,
	SetCount:   3,
	RepCount:   10,
}

var dummyWorkoutExercise = db.Exercise{
	ID:       12345,
	Name:     "some_name",
	Category: "some_category",
}

var dummyWorkoutUser = db.User{
	ID: 12345,
	Name: db.Name{
		FirstName: "some_fname",
		LastName:  "some_lname",
	},
	Age:    33,
	Weight: 205.0,
}

var dummyWorkoutRoutine = db.Routine{
	RoutineID:   5555,
	UserID:      dummyRoutineUser.ID,
	Name:        "some_name",
	Description: "some_description",
	Day:         1,
}

// Workout Dependencies:
// > Exercises (exercise_id)
// > Routines (routine_id) > Users (user_id)
// NOTE: Workout is dependent on both Exercise (exercise_id) and
// Routines (routine_id) which is also dependent on (user_id)
// Implement tests accordingly

func TestGetWorkouts(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/workouts", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts", app.getWorkouts)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddWorkout(t *testing.T) {
	app := setupApp()

	// Add dependencies
	db.AddUser(dummyWorkoutUser, app.Database)
	db.AddRoutine(dummyWorkoutRoutine, app.Database)
	db.AddExercise(dummyWorkoutExercise, app.Database)

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	payload, _ := json.Marshal(dummyWorkout)
	req, err := http.NewRequest("POST", "/workouts", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts", app.addWorkout)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verfiy Workout has been added
	routineid := strconv.Itoa(dummyWorkout.RoutineID)
	exerciseid := strconv.Itoa(dummyWorkout.ExerciseID)
	resWorkout, err := db.GetWorkoutByPKIDs(routineid, exerciseid, app.Database)
	assert.Equal(t, resWorkout, dummyWorkout)
	assert.NoError(t, err)

	// Delete Workout
	db.DeleteWorkoutByPKIDs(routineid, exerciseid, app.Database)

	// Delete dependencies
	db.DeleteExerciseByID(strconv.Itoa(dummyWorkoutExercise.ID), app.Database)
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyWorkoutRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyWorkoutUser.ID), app.Database)
}

func TestGetWorkoutByValidExerciseID(t *testing.T) {
	app := setupApp()

	// Add dependencies
	db.AddUser(dummyWorkoutUser, app.Database)
	db.AddRoutine(dummyWorkoutRoutine, app.Database)
	db.AddExercise(dummyWorkoutExercise, app.Database)

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Add dummyWorkout
	db.AddWorkout(dummyWorkout, app.Database)

	// Verify workout has been added (API Call)
	url := fmt.Sprintf("/workouts/exerciseid/%s", strconv.Itoa(dummyWorkout.ExerciseID))
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/exerciseid/{exerciseid}", app.getWorkoutsByExerciseID)
	app.Router.ServeHTTP(rr, req)
	expected, _ := json.Marshal(db.Workouts{Workouts: []db.Workout{dummyWorkout}})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())

	// Delete Workout
	db.DeleteWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)

	// Delete dependencies
	db.DeleteExerciseByID(strconv.Itoa(dummyWorkoutExercise.ID), app.Database)
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyWorkoutRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyWorkoutUser.ID), app.Database)
}

func TestGetWorkoutByInvalidExerciseID(t *testing.T) {
	app := setupApp()

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Verify workout doesn't exist (API call)
	url := fmt.Sprintf("/workouts/exerciseid/%d", dummyWorkout.ExerciseID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/exerciseid/{exerciseid}", app.getWorkoutsByExerciseID)
	app.Router.ServeHTTP(rr, req)
	expected, _ := json.Marshal(db.Workouts{})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
}

func TestDeleteWorkoutByExerciseID(t *testing.T) {
	app := setupApp()

	// Add dependencies
	db.AddUser(dummyWorkoutUser, app.Database)
	db.AddRoutine(dummyWorkoutRoutine, app.Database)
	db.AddExercise(dummyWorkoutExercise, app.Database)

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Add dummyWorkout
	db.AddWorkout(dummyWorkout, app.Database)

	// Verify it does exist
	result, _ = db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, dummyWorkout)

	// Delete Workout
	url := fmt.Sprintf("/workouts/exerciseid/%d", dummyWorkout.ExerciseID)
	req, err := http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/exerciseid/{exerciseid}", app.deleteWorkoutsByExerciseID)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify workout doesn't exist
	result, _ = db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Delete dependencies
	db.DeleteExerciseByID(strconv.Itoa(dummyWorkoutExercise.ID), app.Database)
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyWorkoutRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyWorkoutUser.ID), app.Database)
}

func TestGetWorkoutByValidRoutineID(t *testing.T) {
	app := setupApp()

	// Add dependencies
	db.AddUser(dummyWorkoutUser, app.Database)
	db.AddRoutine(dummyWorkoutRoutine, app.Database)
	db.AddExercise(dummyWorkoutExercise, app.Database)

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Add dummyWorkout
	db.AddWorkout(dummyWorkout, app.Database)

	// Verify workout has been added (API Call)
	url := fmt.Sprintf("/workouts/routineid/%s", strconv.Itoa(dummyWorkout.RoutineID))
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/routineid/{routineid}", app.getWorkoutsByRoutineID)
	app.Router.ServeHTTP(rr, req)
	expected, _ := json.Marshal(db.Workouts{Workouts: []db.Workout{dummyWorkout}})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())

	// Delete Workout
	db.DeleteWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)

	// Delete dependencies
	db.DeleteExerciseByID(strconv.Itoa(dummyWorkoutExercise.ID), app.Database)
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyWorkoutRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyWorkoutUser.ID), app.Database)
}

func TestGetWorkoutByInvalidRoutineID(t *testing.T) {
	app := setupApp()

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Verify workout doesn't exist (API call)
	url := fmt.Sprintf("/workouts/routineid/%d", dummyWorkout.RoutineID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/routineid/{routineid}", app.getWorkoutsByRoutineID)
	app.Router.ServeHTTP(rr, req)
	expected, _ := json.Marshal(db.Workouts{})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
}

func TestDeleteWorkoutByRoutineID(t *testing.T) {
	app := setupApp()

	// Add dependencies
	db.AddUser(dummyWorkoutUser, app.Database)
	db.AddRoutine(dummyWorkoutRoutine, app.Database)
	db.AddExercise(dummyWorkoutExercise, app.Database)

	// First verify workout doesnt exist yet
	result, _ := db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Add dummyWorkout
	db.AddWorkout(dummyWorkout, app.Database)

	// Verify it does exist
	result, _ = db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, dummyWorkout)

	// Delete Workout
	url := fmt.Sprintf("/workouts/routineid/%d", dummyWorkout.RoutineID)
	req, err := http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/routineid/{routineid}", app.deleteWorkoutsByRoutineID)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify workout doesn't exist
	result, _ = db.GetWorkoutByPKIDs(strconv.Itoa(dummyWorkout.RoutineID), strconv.Itoa(dummyWorkout.ExerciseID), app.Database)
	assert.Equal(t, result, db.Workout{})

	// Delete dependencies
	db.DeleteExerciseByID(strconv.Itoa(dummyWorkoutExercise.ID), app.Database)
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyWorkoutRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyWorkoutUser.ID), app.Database)
}
