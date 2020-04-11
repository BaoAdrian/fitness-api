package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/BaoAdrian/fitness-api/api/db"

	"github.com/stretchr/testify/assert"
)

var dummyExercise = db.Exercise{
	ID:       12345,
	Name:     "some_name",
	Category: "some_category",
}

func TestGetExercises(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises", app.getExercises)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostExercise(t *testing.T) {
	app := setupApp()

	// First verify exercise doesn't exist yet
	result, _ := db.GetExerciseByID(strconv.Itoa(dummyExercise.ID), app.Database)
	assert.Equal(t, result, db.Exercise{})

	// Post exercise
	payload, _ := json.Marshal(dummyExercise)
	req, err := http.NewRequest("POST", "/exercises", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises", app.addExercise)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify exercise exists
	result, err = db.GetExerciseByID(strconv.Itoa(dummyExercise.ID), app.Database)
	assert.Equal(t, result, dummyExercise)
	assert.NoError(t, err)

	// Delete dummy exercise
	db.DeleteExerciseByID(strconv.Itoa(dummyExercise.ID), app.Database)
}

func TestGetExerciseNames(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/names", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/names", app.getExerciseNames)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseCategories(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/categories", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/categories", app.getExerciseCategories)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByValidID(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/id/345", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/id/{exerciseid}", app.getExerciseByID)
	app.Router.ServeHTTP(rr, req)

	expected := `{"exerciseid":345,"name":"hanging pike","category":"abdominals","description":null}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByInvalidID(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/id/99999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/id/{exerciseid}", app.getExerciseByID)
	app.Router.ServeHTTP(rr, req)

	expected := `{"exerciseid":0,"name":"","category":"","description":null}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteExerciseByID(t *testing.T) {
	app := setupApp()

	// First verify exercise doesn't exist
	result, _ := db.GetExerciseByID(strconv.Itoa(dummyExercise.ID), app.Database)
	assert.Equal(t, result, db.Exercise{})

	// Add dummy exercise
	db.AddExercise(dummyExercise, app.Database)

	// Verify exercise exists
	result, _ = db.GetExerciseByID(strconv.Itoa(dummyExercise.ID), app.Database)
	assert.Equal(t, result, dummyExercise)

	// Now Delete the exercise matching ID
	req, err := http.NewRequest("DELETE", "/exercises/id/12345", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/id/{exerciseid}", app.deleteExerciseByID)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify it no longer exists
	result, _ = db.GetExerciseByID(strconv.Itoa(dummyExercise.ID), app.Database)
	assert.Equal(t, result, db.Exercise{})
}

func TestGetExerciseByValidName(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/name/axle%20deadlift", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/name/{name}", app.getExerciseByName)
	app.Router.ServeHTTP(rr, req)

	expected := `{"exerciseid":27,"name":"axle deadlift","category":"lower back","description":null}` + "\n"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetExerciseByInvalidName(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/name/notavalidexercise", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/name/{name}", app.getExerciseByName)
	app.Router.ServeHTTP(rr, req)

	expected := `{"exerciseid":0,"name":"","category":"","description":null}` + "\n"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestDeleteExerciseByName(t *testing.T) {
	app := setupApp()

	// First verify exercise doesn't exist
	result, _ := db.GetExerciseByName(dummyExercise.Name, app.Database)
	assert.Equal(t, result, db.Exercise{})

	// Add dummy exercise
	db.AddExercise(dummyExercise, app.Database)

	// Verify exercise exists
	result, _ = db.GetExerciseByName(dummyExercise.Name, app.Database)
	assert.Equal(t, result, dummyExercise)

	// Now Delete the exercise matching ID
	req, err := http.NewRequest("DELETE", "/exercises/name/some_name", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/name/{name}", app.deleteExerciseByName)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify it no longer exists
	result, _ = db.GetExerciseByName(dummyExercise.Name, app.Database)
	assert.Equal(t, result, db.Exercise{})
}

func TestGetExerciseByValidCategory(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/category/biceps", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/category/{category}", app.getExerciseByCategory)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByInvalidCategory(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/exercises/category/notavalidcategory", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/category/{category}", app.getExerciseByCategory)
	app.Router.ServeHTTP(rr, req)

	expected := `{"exercises":null}` + "\n"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}
