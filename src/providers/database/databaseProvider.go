package databaseProvider

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Shared DatabaseProvider = databaseProvider{}

type DatabaseProvider interface {
	SetupNormalEnvironment()
	SetupLocalEnvironment()

	Select(query string) *sql.Row
	SelectMultiple(query string) (*sql.Rows, error)
	Insert(query string) error
}

type databaseProvider struct {
	db      *sql.DB
	dbError error
}

func (dp databaseProvider) SetupNormalEnvironment() {
	dp.db, dp.dbError = sql.Open("postgres", "host=db port=5432 dbname=test_db user=admin password=test sslmode=disable")
	dp.checkDatabaseReliability()
}

func (dp databaseProvider) SetupLocalEnvironment() {
	dp.db, dp.dbError = sql.Open("postgres", "host=localhost port=5432 dbname=test_db user=admin password=test sslmode=disable")
	dp.checkDatabaseReliability()
}

// Insert implements DatabaseProvider.
func (dp databaseProvider) Insert(query string) error {
	if dp.db == nil {
		panic("database not instantiated")
	}

	_, queryErr := dp.db.Exec(query)

	dp.checkDatabaseReliability()

	if queryErr != nil {
		return queryErr
	}

	return nil
}

// Select implements DatabaseProvider.
func (dp databaseProvider) Select(query string) *sql.Row {
	if dp.db == nil {
		panic("database not instantiated")
	}

	dp.checkDatabaseReliability()
	return dp.db.QueryRow(query)
}

// SelectMultiple implements DatabaseProvider.
func (dp databaseProvider) SelectMultiple(query string) (*sql.Rows, error) {
	if dp.db == nil {
		panic("database not instantiated")
	}

	dp.checkDatabaseReliability()
	return dp.db.Query(query)
}

func (dp databaseProvider) checkDatabaseReliability() {
	if dp.db.Ping() != nil {
		dp.db.Close()
		panic("couldn't instantiate database connection!!!")
	}
}
