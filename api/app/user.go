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

// (GET) Endpoint: /users
func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	users, err := db.GetUsers(app.Database)
	if err = json.NewEncoder(w).Encode(users); err != nil {
		log.Fatal("ERROR: Failed to Encode Users")
	}
}

// (POST) Endpoint: /users
func (app *App) addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal("ERROR: Failed to read request body.")
	}
	var user db.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal("ERROR: Failed to unpack user data. Verify JSON formatting.")
	}

	db.AddUser(user, app.Database)
}

// (GET) Endpoint: /users/id/{userid}
func (app *App) getUserByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["userid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	results, err := db.GetUserByUserID(id, app.Database)
	if err = json.NewEncoder(w).Encode(results); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /users/id/{userid}
func (app *App) deleteUserByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["userid"]

	if !ok {
		log.Fatal("ERROR: No id was provided")
	}

	db.DeleteUserByUserID(id, app.Database)
}
