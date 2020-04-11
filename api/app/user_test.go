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

var dummyUser = db.User{
	ID: 12345,
	Name: db.Name{
		FirstName: "some_fname",
		LastName:  "some_lname",
	},
	Age:    33,
	Weight: 205.0,
}

func TestGetUsers(t *testing.T) {
	app := setupApp()

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/users", app.getUsers)
	app.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAddUser(t *testing.T) {
	app := setupApp()

	// Verify User doesn't exist yet
	result, _ := db.GetUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
	assert.Equal(t, result, db.User{})

	// Add user
	payload, _ := json.Marshal(dummyUser)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/users", app.addUser)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verify User exists
	result, err = db.GetUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
	assert.Equal(t, result, dummyUser)
	assert.NoError(t, err)

	// Delete dummy exercise
	db.DeleteUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
}

func TestGetUserByValidUserID(t *testing.T) {
	app := setupApp()

	// Verify User doesn't exist yet (using direct DB call)
	result, _ := db.GetUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
	assert.Equal(t, result, db.User{})

	// Add user
	db.AddUser(dummyUser, app.Database)

	// GET User by user id
	req, err := http.NewRequest("GET", "/users/id/12345", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/users/id/{userid}", app.getUserByUserID)
	app.Router.ServeHTTP(rr, req)

	expected, _ := json.Marshal(dummyUser)
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)

	// Delete User
	db.DeleteUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
}

func TestGetUserByInvalidUserID(t *testing.T) {
	app := setupApp()

	invalidID := "99999"

	// Verify User doesn't exist yet (using direct DB call)
	result, _ := db.GetUserByUserID(invalidID, app.Database)
	assert.Equal(t, result, db.User{})

	// GET User by invalid user, hitting endpoint
	req, err := http.NewRequest("GET", "/users/id/99999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/users/id/{userid}", app.getUserByUserID)
	app.Router.ServeHTTP(rr, req)

	expected, _ := json.Marshal(db.User{})
	assert.Equal(t, string(expected)+"\n", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteUserByUserID(t *testing.T) {
	app := setupApp()

	// Verify User doesn't exist yet (using direct DB call)
	result, _ := db.GetUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
	assert.Equal(t, result, db.User{})

	// Add user
	db.AddUser(dummyUser, app.Database)

	// Verify User now exists
	result, _ = db.GetUserByUserID(strconv.Itoa(dummyUser.ID), app.Database)
	assert.Equal(t, result, dummyUser)

	// Delete User
	req, err := http.NewRequest("DELETE", "/users/id/99999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	app.Router.HandleFunc("/users/id/{userid}", app.deleteUserByUserID)
	app.Router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
