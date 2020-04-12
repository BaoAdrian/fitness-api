package db

import (
	"database/sql"
	"fmt"
	"log"
)

// Workout Struct
type Workout struct {
	ExerciseID int `json:"exerciseid"`
	RoutineID  int `json:"routineid"`
	SetCount   int `json:"setcount"`
	RepCount   int `json:"repcount"`
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
		if err = rows.Scan(&workout.ExerciseID, &workout.RoutineID, &workout.SetCount, &workout.RepCount); err != nil {
			log.Fatal("ERROR: Failed to Parse Workouts")
		}
		collection.Workouts = append(collection.Workouts, workout)
	}

	return collection, err
}

// AddWorkout Adds given workout to the database
func AddWorkout(workout Workout, db *sql.DB) {
	query := fmt.Sprintf(`INSERT INTO Workouts (exercise_id, routine_id, set_count, rep_count) VALUES (%d,%d,%d,%d)`,
		workout.ExerciseID, workout.RoutineID, workout.SetCount, workout.RepCount)
	db.Exec(query)
}

// GetWorkoutsByExerciseID Retrieves workouts associated with a specific exercise (given by exercise_id)
func GetWorkoutsByExerciseID(id string, db *sql.DB) (workouts Workouts, err error) {
	// Gather all records matching id
	rows, _ := runQuery(db, fmt.Sprintf("SELECT * FROM Workouts WHERE exercise_id = %s", id))

	collection := Workouts{}
	for rows.Next() {
		workout := Workout{}
		if err = rows.Scan(&workout.ExerciseID, &workout.RoutineID, &workout.SetCount, &workout.RepCount); err != nil {
			log.Fatal("SELECT Failed to Retrieve Workouts for id: " + id)
		}
		collection.Workouts = append(collection.Workouts, workout)
	}
	return collection, err
}

// DeleteWorkoutByExerciseID Removes workout matching provided 'exercise_id'
func DeleteWorkoutByExerciseID(id string, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM Workouts WHERE exercise_id = %s", id)
	db.Exec(query)
}

// GetWorkoutsByRoutineID Retrieves workouts matching the given 'routine_id'
func GetWorkoutsByRoutineID(id string, db *sql.DB) (workouts Workouts, err error) {
	rows, err := runQuery(db, fmt.Sprintf(`SELECT * FROM Workouts WHERE routine_id = "%s"`, id))
	if err != nil {
		log.Fatal("ERROR: SELECT Failed to Retrieve Workouts matching id: " + id)
	}

	collection := Workouts{}
	for rows.Next() {
		workout := Workout{}
		if err = rows.Scan(&workout.ExerciseID, &workout.RoutineID, &workout.SetCount, &workout.RepCount); err != nil {
			log.Fatal("SELECT Failed to Retrieve Workouts for id: " + id)
		}
		collection.Workouts = append(collection.Workouts, workout)
	}
	return collection, err
}

// DeleteWorkoutByRoutineID Deletes workouts matching the given 'routine_id'
func DeleteWorkoutByRoutineID(id string, db *sql.DB) {
	query := fmt.Sprintf(`DELETE FROM Workouts WHERE routine_id = "%s"`, id)
	db.Exec(query)
}

// GetWorkoutByPKIDs Retrieves the Workout associated with a given exerciseid and routineid
func GetWorkoutByPKIDs(routineid string, exerciseid string, db *sql.DB) (workout Workout, err error) {
	query := fmt.Sprintf(`SELECT * FROM Workouts WHERE routine_id = "%s" AND exercise_id = "%s"`, routineid, exerciseid)
	result := Workout{}
	err = db.QueryRow(query).Scan(&result.ExerciseID, &result.RoutineID, &result.SetCount, &result.RepCount)
	return result, err
}

// DeleteWorkoutByPKIDs Removes Workout with matching (routineid, exerciseid) PK
func DeleteWorkoutByPKIDs(routineid string, exerciseid string, db *sql.DB) {
	query := fmt.Sprintf(`DELETE FROM Workouts WHERE routine_id = "%s" AND exercise_id = "%s"`, routineid, exerciseid)
	db.Exec(query)
}
