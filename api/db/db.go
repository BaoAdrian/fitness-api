package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Exercise struct
type Exercise struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description *string `json:"description"`
}

// Exercises Struct
type Exercises struct {
	Exercises []Exercise `json:"exercises"`
}

// Names Struct
type Names struct {
	Names []string `json:"names"`
}

// Category Struct
type Category struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// Categories struct
type Categories struct {
	Categories []Category `json:"categories"`
}

// CreateDatabase Function
func CreateDatabase(host string) (*sql.DB, error) {
	user := "root"
	pass := "password"
	dbName := "fitnessdb"

	source := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbName)

	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetExercises Retrieves All Exercises from Database
func GetExercises(db *sql.DB) (exercises Exercises, err error) {
	rows, err := runQuery(db, `SELECT * FROM exercises`)
	if err != nil {
		log.Fatal("runQuery Failed")
	}

	collection := Exercises{}
	for rows.Next() {
		exercise := Exercise{}
		if err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Exercises = append(collection.Exercises, exercise)
	}

	return collection, err
}

// GetExerciseNames Retrieves all exercise names from database
func GetExerciseNames(db *sql.DB) (names Names, err error) {
	rows, err := runQuery(db, `SELECT name FROM exercises`)

	collection := Names{}
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Names = append(collection.Names, name)
	}

	return collection, err
}

// GetExerciseCategories Retrieves all exercise categories
func GetExerciseCategories(db *sql.DB) (categories Categories, err error) {
	rows, err := runQuery(db, `SELECT category, COUNT(*) FROM exercises GROUP BY category`)

	collection := struct {
		Categories []Category `json:"categories"`
	}{}
	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.Category, &category.Count); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Categories = append(collection.Categories, category)
	}

	return collection, err
}

// GetExerciseByID Retrieves exercise matching the given id
func GetExerciseByID(id string, db *sql.DB) (exercise Exercise, err error) {
	result := Exercise{}
	err = db.QueryRow(fmt.Sprintf("SELECT * FROM exercises WHERE exerciseid = %s", id)).Scan(&result.ID, &result.Name, &result.Category, &result.Description)
	return result, err
}

// GetExerciseByName Retrieves exercise matching the given name
func GetExerciseByName(name string, db *sql.DB) (exercise Exercise, err error) {
	result := Exercise{}
	err = db.QueryRow(fmt.Sprintf(`SELECT * FROM exercises WHERE name = "%s"`, name)).Scan(&result.ID, &result.Name, &result.Category, &result.Description)
	return result, err
}

// GetExerciseByCategory Retrieves exercise(s) matching given category
func GetExerciseByCategory(category string, db *sql.DB) (exercises Exercises, err error) {
	rows, err := runQuery(db, fmt.Sprintf(`SELECT * FROM exercises WHERE category = "%s"`, category))

	collection := Exercises{}
	for rows.Next() {
		exercise := Exercise{}
		if err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Category, &exercise.Description); err != nil {
			log.Fatal("Database SELECT failed")
		}
		collection.Exercises = append(collection.Exercises, exercise)
	}
	return collection, err
}

// AddExercise Adds exercise to database
func AddExercise(exercise Exercise, db *sql.DB) {
	var query string
	if exercise.Description != nil {
		query = fmt.Sprintf(`INSERT INTO exercises (exerciseid,name,category,description) VALUES (%d,"%s","%s","%s")`, exercise.ID, exercise.Name, exercise.Category, *exercise.Description)
	} else {
		query = fmt.Sprintf(`INSERT INTO exercises (exerciseid,name,category) VALUES (%d,"%s","%s")`, exercise.ID, exercise.Name, exercise.Category)
	}
	db.Exec(query)
}

// DeleteExerciseByID Deletes exercise with a given exerciseid
func DeleteExerciseByID(exerciseid string, db *sql.DB) {
	query := fmt.Sprintf("DELETE FROM exercises WHERE exerciseid = %s", exerciseid)
	db.Exec(query)
}

// DeleteExerciseByName Deletes exercise with a given name
func DeleteExerciseByName(name string, db *sql.DB) {
	query := fmt.Sprintf(`DELETE FROM exercises WHERE name = "%s"`, name)
	db.Exec(query)
}

// RunQuery Runs specified query on database
func runQuery(db *sql.DB, query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
