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

// Note: Routines is dependant (FK) on a User
var dummyRoutineUser = db.User{
	ID: 12345,
	Name: db.Name{
		FirstName: "some_fname",
		LastName:  "some_lname",
	},
	Age:    33,
	Weight: 205.0,
}

var dummyRoutine = db.Routine{
	RoutineID:   5555,
	UserID:      dummyRoutineUser.ID,
	Name:        "some_name",
	Description: "some_description",
	Day:         1,
}

var anotherDummyRoutine = db.Routine{
	RoutineID:   77777,
	UserID:      dummyRoutineUser.ID,
	Name:        "some_other_name",
	Description: "some_other_description",
	Day:         2,
}

func TestGetRoutines(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/routines", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines", app.getUsers)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddRoutine(t *testing.T) {
	app := setupApp()

	// First add user to User table (dependancy)
	db.AddUser(dummyRoutineUser, app.Database)

	// Verify routine doesn't exist with given id (direct DB call)
	result, _ := db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})

	// POST Routine using
	payload, _ := json.Marshal(dummyRoutine)
	req, err := http.NewRequest("POST", "/routines", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines", app.addRoutine)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify routine exists
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, dummyRoutine)

	// Delete Routine & User
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyRoutineUser.ID), app.Database)
}

func TestGetRoutineByValidRoutineID(t *testing.T) {
	app := setupApp()

	// First add user to User table (dependancy)
	db.AddUser(dummyRoutineUser, app.Database)

	// Verify routine doesn't exist with given id (direct DB call)
	result, _ := db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})

	// Add Routine
	db.AddRoutine(dummyRoutine, app.Database)

	// GET Routine using routineid (API request)
	url := fmt.Sprintf("/routines/routineid/%d", dummyRoutine.RoutineID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines/routineid/{routineid}", app.getRoutineByRoutineID)
	app.Router.ServeHTTP(rr, req)

	expected, _ := json.Marshal(dummyRoutine)
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)

	// Delete Routine & User
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyRoutineUser.ID), app.Database)
}

func TestGetRoutineByInvalidRoutineID(t *testing.T) {
	app := setupApp()

	invalidID := "99999"

	// Verify routine doesn't exist with given id (direct DB call)
	result, _ := db.GetRoutineByRoutineID(invalidID, app.Database)
	assert.Equal(t, result, db.Routine{})

	// GET Routine using routineid (API request)
	url := fmt.Sprintf("/routines/routineid/%d", dummyRoutine.RoutineID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines/routineid/{routineid}", app.getRoutineByRoutineID)
	app.Router.ServeHTTP(rr, req)

	expected, _ := json.Marshal(db.Routine{})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteRoutineByRoutineID(t *testing.T) {
	app := setupApp()

	// First add user to User table (dependancy)
	db.AddUser(dummyRoutineUser, app.Database)

	// Verify Routine doesn't exist
	result, _ := db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})

	// Add Routine
	db.AddRoutine(dummyRoutine, app.Database)

	// Verify it now exists
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, dummyRoutine)

	// DELETE Routine (API Call)
	url := fmt.Sprintf("/routines/routineid/%d", dummyRoutine.RoutineID)
	req, err := http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines/routineid/{routineid}", app.deleteRoutineByRoutineID)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Delete User
	db.DeleteUserByUserID(strconv.Itoa(dummyRoutineUser.ID), app.Database)
}

func TestGetRoutineByValidUserID(t *testing.T) {
	app := setupApp()

	// First add user to User table (dependancy)
	db.AddUser(dummyRoutineUser, app.Database)

	// Verify routine(s) doesn't exist with given id (direct DB call)
	result, _ := db.GetRoutinesByUserID(strconv.Itoa(dummyRoutine.UserID), app.Database)
	assert.Equal(t, result, db.Routines{})

	// Add Some Routines
	db.AddRoutine(dummyRoutine, app.Database)
	db.AddRoutine(anotherDummyRoutine, app.Database)

	// GET Routines using userid (API request)
	url := fmt.Sprintf("/routines/userid/%d", dummyRoutineUser.ID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines/userid/{userid}", app.getRoutinesByUserID)
	app.Router.ServeHTTP(rr, req)

	routines := db.Routines{}
	routines.Routines = append(routines.Routines, dummyRoutine)
	routines.Routines = append(routines.Routines, anotherDummyRoutine)

	expected, _ := json.Marshal(routines)
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)

	// Delete Routines & User
	db.DeleteRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	db.DeleteRoutineByRoutineID(strconv.Itoa(anotherDummyRoutine.RoutineID), app.Database)
	db.DeleteUserByUserID(strconv.Itoa(dummyRoutineUser.ID), app.Database)
}

func TestGetRoutineByInvalidUserID(t *testing.T) {
	app := setupApp()

	// Verify routine(s) don't exist with given id (direct DB call)
	result, _ := db.GetRoutinesByUserID(strconv.Itoa(dummyRoutine.UserID), app.Database)
	assert.Equal(t, result, db.Routines{})

	// GET User by invalid user, hitting endpoint
	url := fmt.Sprintf("/routines/userid/%d", dummyRoutineUser.ID)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines/userid/{userid}", app.getRoutinesByUserID)
	app.Router.ServeHTTP(rr, req)

	expected, _ := json.Marshal(db.Routines{})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeletRoutinesByUserID(t *testing.T) {
	app := setupApp()

	// First add user to User table (dependancy)
	db.AddUser(dummyRoutineUser, app.Database)

	// Verify Routines dont't exist
	result, _ := db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(anotherDummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})

	// Add Routines
	db.AddRoutine(dummyRoutine, app.Database)
	db.AddRoutine(anotherDummyRoutine, app.Database)

	// Verify they now exists
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, dummyRoutine)
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(anotherDummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, anotherDummyRoutine)

	// DELETE Routines (API Call)
	url := fmt.Sprintf("/routines/userid/%d", dummyRoutineUser.ID)
	req, err := http.NewRequest("DELETE", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/routines/userid/{userid}", app.deleteRoutinesByUserID)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify BOTH Routines have been deleted
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(dummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})
	result, _ = db.GetRoutineByRoutineID(strconv.Itoa(anotherDummyRoutine.RoutineID), app.Database)
	assert.Equal(t, result, db.Routine{})

	// Delete User
	db.DeleteUserByUserID(strconv.Itoa(dummyRoutineUser.ID), app.Database)
}
