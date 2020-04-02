package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BaoAdrian/fitness-api/api/db"

	"github.com/stretchr/testify/assert"
)

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

	// First verifiy workout doesnt exist yet
	result, _ := db.GetWorkoutByWorkoutID("12345", app.Database)
	assert.Equal(t, result, db.Workouts{})

	// Post Workout
	dummyWorkout := db.Workout{
		ID:         12345,
		Name:       "some_name",
		ExerciseID: 1,
		SetCount:   4,
		RepCount:   10,
	}
	payload, _ := json.Marshal(dummyWorkout)
	req, err := http.NewRequest("POST", "/workouts", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts", app.addWorkout)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verfiy Workout has been added
	expected := db.Workouts{Workouts: []db.Workout{dummyWorkout}}
	result, err = db.GetWorkoutByWorkoutID("12345", app.Database)
	assert.Equal(t, result, expected)
	assert.NoError(t, err)

	// Delete Workout
	db.DeleteWorkoutByID("12345", app.Database)
}

func TestGetWorkoutsByValidWorkoutID(t *testing.T) {
	app := setupApp()

	// Verify workout doesn't exist yet
	result, _ := db.GetWorkoutByWorkoutID("12345", app.Database)
	assert.Equal(t, result, db.Workouts{})

	// Add dummy workout
	dummyWorkout := db.Workout{
		ID:         12345,
		Name:       "some_name",
		ExerciseID: 1,
		SetCount:   4,
		RepCount:   10,
	}
	db.AddWorkout(dummyWorkout, app.Database)

	// GET workout by workoutid
	req, err := http.NewRequest("GET", "/workouts/id/12345", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/id/{workoutid}", app.getWorkoutByWorkoutID)
	app.Router.ServeHTTP(rr, req)

	expected := `{"workouts":[{"workoutid":12345,"name":"some_name","exerciseid":1,"setcount":4,"repcount":10}]}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)

	// Delete workout
	db.DeleteWorkoutByID("12345", app.Database)
}

func TestGetWorkoutsByInvalidWorkoutID(t *testing.T) {
	app := setupApp()

	// Verify workout doesn't exist (using direct query as opposed request)
	result, _ := db.GetWorkoutByWorkoutID("12345", app.Database)
	assert.Equal(t, result, db.Workouts{})

	// Test API Endpoint to verify workout doesn't exist
	req, err := http.NewRequest("GET", "/workouts/id/12345", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/id/{workoutid}", app.getWorkoutByWorkoutID)
	app.Router.ServeHTTP(rr, req)

	expected := `{"workouts":null}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetWorkoutsByValidName(t *testing.T) {
	app := setupApp()

	// Verify workout doesn't exist yet
	result, _ := db.GetWorkoutByName("some_name", app.Database)
	assert.Equal(t, result, db.Workouts{})

	// Add dummy workout
	dummyWorkout := db.Workout{
		ID:         12345,
		Name:       "some_name",
		ExerciseID: 1,
		SetCount:   4,
		RepCount:   10,
	}
	db.AddWorkout(dummyWorkout, app.Database)

	// GET workout by workoutid
	req, err := http.NewRequest("GET", "/workouts/name/some_name", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/workouts/name/{name}", app.getWorkoutByWorkoutName)
	app.Router.ServeHTTP(rr, req)

	expected := `{"workouts":[{"workoutid":12345,"name":"some_name","exerciseid":1,"setcount":4,"repcount":10}]}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)

	// Delete workout
	db.DeleteWorkoutByName("some_name", app.Database)
}
