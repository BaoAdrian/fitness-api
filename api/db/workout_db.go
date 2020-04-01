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
		if err = rows.Scan(&workout.ID, &workout.Name, &workout.ExerciseID); err != nil {
			log.Fatal("ERROR: Failed to Parse Workouts")
		}
		collection.Workouts = append(collection.Workouts, workout)
	}

	return collection, err
}

// AddWorkout Adds given workout to the database
func AddWorkout(workout Workout, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO Workouts (workoutid,name,exerciseid) VALUES (%d,"%s","%d")`, workout.ID, workout.Name, workout.ExerciseID)
	db.Exec(query)
}
