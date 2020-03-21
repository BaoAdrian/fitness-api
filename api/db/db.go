package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	host := "db:3306"
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

// RunQuery Runs specified query on database
func RunQuery(db *sql.DB, query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
