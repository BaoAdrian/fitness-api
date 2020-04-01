package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
