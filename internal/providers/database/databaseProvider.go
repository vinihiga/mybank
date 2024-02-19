package databaseProvider

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type IDatabaseProvider interface {
	SetupNormalEnvironment()
	SetupLocalEnvironment()

	Select(query string, args ...any) (*sql.Rows, error)
	Insert(query string, args ...any) error
}

type DatabaseProvider struct {
	db      *sql.DB
	dbError error
}

func (dp *DatabaseProvider) SetupNormalEnvironment() {
	dp.db, dp.dbError = sql.Open("postgres", "host=db port=5432 dbname=test_db user=admin password=test sslmode=disable")
	dp.checkDatabaseReliability()
}

func (dp *DatabaseProvider) SetupLocalEnvironment() {
	dp.db, dp.dbError = sql.Open("postgres", "host=localhost port=5432 dbname=test_db user=admin password=test sslmode=disable")
	dp.checkDatabaseReliability()
}

// Insert implements DatabaseProvider.
func (dp *DatabaseProvider) Insert(query string, args ...any) error {
	dp.checkDatabaseReliability()
	_, queryErr := dp.db.Exec(query, args...)

	if queryErr != nil {
		return queryErr
	}

	return nil
}

// Select implements DatabaseProvider.
func (dp *DatabaseProvider) Select(query string, args ...any) (*sql.Rows, error) {
	dp.checkDatabaseReliability()
	return dp.db.Query(query, args...)
}

func (dp *DatabaseProvider) checkDatabaseReliability() {
	if dp.db == nil {
		panic("database not instantiated")
	} else if dp.db.Ping() != nil {
		dp.db.Close()
		panic("couldn't instantiate database connection!!!")
	}
}
