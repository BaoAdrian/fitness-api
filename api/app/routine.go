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

// (GET) Endpoint: /routines
func (app *App) getRoutines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	users, err := db.GetRoutines(app.Database)
	if err = json.NewEncoder(w).Encode(users); err != nil {
		log.Fatal("ERROR: Failed to Encode Routines")
	}
}

// (POST) Endpoint: /routines
func (app *App) addRoutine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal("ERROR: Failed to read request body.")
	}
	var routine db.Routine
	if err := json.Unmarshal(body, &routine); err != nil {
		log.Fatal("ERROR: Failed to unpack routine data. Verify JSON formatting.")
	}

	db.AddRoutine(routine, app.Database)
}

// (GET) Endpoint: /routines/routineid/{routineid}
func (app *App) getRoutineByRoutineID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["routineid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	results, err := db.GetRoutineByRoutineID(id, app.Database)
	if err = json.NewEncoder(w).Encode(results); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /routines/routineid/{routineid}
func (app *App) deleteRoutineByRoutineID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["routineid"]

	if !ok {
		log.Fatal("ERROR: No id was provided")
	}

	db.DeleteRoutineByRoutineID(id, app.Database)
}

// (GET) Endpoint: /routines/userid/{userid}
func (app *App) getRoutinesByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["userid"]
	if !ok {
		log.Fatal("ERROR: No ID was provided")
	}

	results, err := db.GetRoutinesByUserID(id, app.Database)
	if err = json.NewEncoder(w).Encode(results); err != nil {
		log.Fatal("ERROR: Failed to Encode JSON")
	}
}

// (DELETE) Endpoint: /routines/userid/{userid}
func (app *App) deleteRoutinesByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, ok := vars["userid"]

	if !ok {
		log.Fatal("ERROR: No id was provided")
	}

	db.DeleteRoutinesByUserID(id, app.Database)
}
