package databaseProvider

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type IDatabaseProvider interface {
	SetupNormalEnvironment()
	SetupLocalEnvironment()

	Select(query string) (*sql.Rows, error)
	Insert(query string) error
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
func (dp *DatabaseProvider) Insert(query string) error {
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
func (dp *DatabaseProvider) Select(query string) (*sql.Rows, error) {
	if dp.db == nil {
		panic("database not instantiated")
	}

	dp.checkDatabaseReliability()
	return dp.db.Query(query)
}

func (dp *DatabaseProvider) checkDatabaseReliability() {
	if dp.db.Ping() != nil {
		dp.db.Close()
		panic("couldn't instantiate database connection!!!")
	}
}
