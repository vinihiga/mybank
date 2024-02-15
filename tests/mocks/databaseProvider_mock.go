package mocks

import "database/sql"

/*

	SetupNormalEnvironment()
	SetupLocalEnvironment()

	Select(query string) *sql.Row
	SelectMultiple(query string) (*sql.Rows, error)
	Insert(query string) error

*/

type DatabaseProviderMock struct {
}

func (dp *DatabaseProviderMock) SetupNormalEnvironment() {
	// Intentionally not implemented
}

func (dp *DatabaseProviderMock) SetupLocalEnvironment() {
	// Intentionally not implemented
}

func (dp *DatabaseProviderMock) Select(query string) *sql.Row {
	return nil
}

func (dp *DatabaseProviderMock) SelectMultiple(query string) (*sql.Rows, error) {
	return nil, nil
}

func (dp *DatabaseProviderMock) Insert(query string) error {
	return nil
}
