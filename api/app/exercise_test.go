package app

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BaoAdrian/fitness-api/api/db"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const apiURL = "http://localhost:8080"

// NOTE: Database connection created using 'localhost' hostname
// to resolve tcp unable to connect to 'db' container_name as the
// MySQL host.
func createTestDatabase() (*sql.DB, error) {
	database, err := db.CreateDatabase("localhost:3306")
	return database, err
}

func TestGetExercises(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises", app.getExercises)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseNames(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/names", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/names", app.getExerciseNames)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseCategories(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/categories", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/categories", app.getExerciseCategories)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByValidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/id/345", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/id/{exerciseid}", app.getExerciseByID)
	app.Router.ServeHTTP(rr, req)

	expected := `{"id":345,"name":"hanging pike","category":"abdominals","description":null}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/id/99999", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/id/{exerciseid}", app.getExerciseByID)
	app.Router.ServeHTTP(rr, req)

	expected := `{"message":"No data found."}` + "\n"
	assert.Equal(t, expected, rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByValidName(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/name/axle%20deadlift", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/name/{name}", app.getExerciseByName)
	app.Router.ServeHTTP(rr, req)

	expected := `{"id":27,"name":"axle deadlift","category":"lower back","description":null}` + "\n"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetExerciseByInvalidName(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/name/notavalidexercise", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/name/{name}", app.getExerciseByName)
	app.Router.ServeHTTP(rr, req)

	expected := `{"message":"No data found."}` + "\n"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetExerciseByValidCategory(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/category/biceps", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/category/{category}", app.getExerciseByCategory)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetExerciseByInvalidCategory(t *testing.T) {
	req, err := http.NewRequest("GET", "/exercises/category/notavalidcategory", nil)
	assert.NoError(t, err)

	database, err := createTestDatabase()
	assert.NoError(t, err)

	app := App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/exercises/category/{category}", app.getExerciseByCategory)
	app.Router.ServeHTTP(rr, req)

	expected := `{"exercises":null}` + "\n"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected, rr.Body.String())
}
