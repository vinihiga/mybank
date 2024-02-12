package databaseProvider

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db, dbError = sql.Open("postgres", "host=db port=5432 dbname=test_db user=admin password=test sslmode=disable")

func Select(query string) *sql.Row {
	if dbError != nil {
		return nil
	}

	return db.QueryRow(query)
}

func Insert(query string) error {
	_, queryErr := db.Exec(query)

	if queryErr != nil {
		return queryErr
	}

	return nil
}

func SetupLocalEnvironment() {
	db, dbError = sql.Open("postgres", "host=localhost port=5432 dbname=test_db user=admin password=test sslmode=disable")
}
