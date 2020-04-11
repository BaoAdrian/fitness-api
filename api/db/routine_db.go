package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Routine struct
type Routine struct {
	RoutineID   int    `json:"routineid"`
	UserID      int    `json:"userid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Day         int    `json:"day"`
}

// Routines struct
type Routines struct {
	Routines []Routine `json:"routines"`
}

// GetRoutines Retrieves all routines from the database
func GetRoutines(db *sql.DB) (routines Routines, err error) {
	rows, err := runQuery(db, "SELECT * FROM Routines")
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve Routines")
	}

	collection := Routines{}
	for rows.Next() {
		routine := Routine{}
		if err = rows.Scan(&routine.RoutineID, &routine.UserID, &routine.Name, &routine.Description, &routine.Day); err != nil {
			log.Fatal("ERROR: Failed to Parse User Data")
		}
		collection.Routines = append(collection.Routines, routine)
	}

	return collection, err
}

// AddRoutine Adds given Routine to the database
func AddRoutine(routine Routine, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO Routines (routine_id, user_id, routine_name, description, day) VALUES (%d, %d, "%s", "%s", %d)`,
		routine.RoutineID, routine.UserID, routine.Name, routine.Description, routine.Day)
	db.Exec(query)
}

// GetRoutineByRoutineID Retrieves Routine matching given 'id'
func GetRoutineByRoutineID(id string, db *sql.DB) (routine Routine, err error) {
	resRoutine := Routine{}
	query := fmt.Sprintf(`SELECT * FROM Routines WHERE routine_id = "%s"`, id)
	err = db.QueryRow(query).Scan(&resRoutine.RoutineID, &resRoutine.UserID, &resRoutine.Name, &resRoutine.Description, &resRoutine.Day)
	return resRoutine, err
}

// DeleteRoutineByRoutineID Removes Routine matching given 'id'
func DeleteRoutineByRoutineID(id string, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM Routines WHERE routine_id = %s", id)
	db.Exec(query)
}

// GetRoutinesByUserID Retrieves Routine matching given 'id'
func GetRoutinesByUserID(id string, db *sql.DB) (routines Routines, err error) {
	query := fmt.Sprintf(`SELECT * FROM Routines WHERE user_id = "%s"`, id)
	rows, err := runQuery(db, query)
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve Routines")
	}

	collection := Routines{}
	for rows.Next() {
		routine := Routine{}
		if err = rows.Scan(&routine.RoutineID, &routine.UserID, &routine.Name, &routine.Description, &routine.Day); err != nil {
			log.Fatal("ERROR: Failed to Parse User Data")
		}
		collection.Routines = append(collection.Routines, routine)
	}

	return collection, err
}

// DeleteRoutinesByUserID Removes Routine matching given 'id'
func DeleteRoutinesByUserID(id string, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM Routines WHERE user_id = %s", id)
	db.Exec(query)
}
