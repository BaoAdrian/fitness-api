package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Workout Struct
type Workout struct {
	ID         int    `json:"workoutid"`
	Name       string `json:"name"`
	ExerciseID int    `json:"exerciseid"`
	SetCount   int    `json:"setcount"`
	RepCount   int    `json:"repcount"`
}

// Workouts Struct
type Workouts struct {
	Workouts []Workout `json:"workouts"`
}

// GetWorkouts Retrieves all existing workouts in database
func GetWorkouts(db *sql.DB) (workout Workouts, err error) {
	rows, err := runQuery(db, `SELECT * FROM Workouts`)
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve Workouts")
	}

	collection := Workouts{}
	for rows.Next() {
		workout := Workout{}
		if err = rows.Scan(&workout.ID, &workout.Name, &workout.ExerciseID, &workout.SetCount, &workout.RepCount); err != nil {
			log.Fatal("ERROR: Failed to Parse Workouts")
		}
		collection.Workouts = append(collection.Workouts, workout)
	}

	return collection, err
}

// AddWorkout Adds given workout to the database
func AddWorkout(workout Workout, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO Workouts (workoutid,name,exerciseid,setcount,repcount) VALUES (%d,"%s",%d,%d,%d)`, workout.ID, workout.Name, workout.ExerciseID, workout.SetCount, workout.RepCount)
	db.Exec(query)
}

// GetWorkoutByWorkoutID Retrieves workout matching given 'workoutid'
func GetWorkoutByWorkoutID(id string, db *sql.DB) (workouts Workouts, err error) {
	// Gather all records matching id
	rows, _ := runQuery(db, fmt.Sprintf("SELECT * FROM Workouts WHERE workoutid = %s", id))

	collection := Workouts{}
	for rows.Next() {
		workout := Workout{}
		if err = rows.Scan(&workout.ID, &workout.Name, &workout.ExerciseID, &workout.SetCount, &workout.RepCount); err != nil {
			log.Fatal("SELECT Failed to Retrieve Workouts for id: " + id)
		}
		collection.Workouts = append(collection.Workouts, workout)
	}
	return collection, err
}

// DeleteWorkoutByID Removes workout matching provided 'workoutid'
func DeleteWorkoutByID(id string, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM Workouts WHERE workoutid = %s", id)
	db.Exec(query)
}

// GetWorkoutByName Retrieves workouts matching the given name
func GetWorkoutByName(name string, db *sql.DB) (workouts Workouts, err error) {
	rows, err := runQuery(db, fmt.Sprintf(`SELECT * FROM Workouts WHERE name = "%s"`, name))
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve Workouts matching name: " + name)
	}

	collection := Workouts{}
	for rows.Next() {
		workout := Workout{}
		if err = rows.Scan(&workout.ID, &workout.Name, &workout.ExerciseID, &workout.SetCount, &workout.RepCount); err != nil {
			log.Fatal("SELECT Failed to Retrieve Workouts for name: " + name)
		}
		collection.Workouts = append(collection.Workouts, workout)
	}
	return collection, err
}

// DeleteWorkoutByName Retrieves workouts matching the given name
func DeleteWorkoutByName(name string, db *sql.DB) {
	query := fmt.Sprintf(`DELETE FROM Workouts WHERE name = "%s"`, name)
	db.Exec(query)
}

// GetExerciseIDByWorkoutID Retrieves exerciseids from workoutid
func GetExerciseIDByWorkoutID(id string, db *sql.DB) (results []int) {
	rows, err := runQuery(db, fmt.Sprintf("SELECT exerciseid FROM Workouts WHERE workoutid = %s", id))
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve exerciseids matching workoutid: " + id)
	}

	var exerciseids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal("SELECT Failed to Parse exerciseids")
		}
		exerciseids = append(exerciseids, id)
	}
	return exerciseids
}
