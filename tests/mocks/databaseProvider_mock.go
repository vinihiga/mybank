package mocks

import (
	"database/sql"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
)

type DatabaseProviderMock struct{}

func (dp *DatabaseProviderMock) SetupNormalEnvironment() {
	// Intentionally not implemented
}

func (dp *DatabaseProviderMock) SetupLocalEnvironment() {
	// Intentionally not implemented
}

func (dp *DatabaseProviderMock) Select(query string, args ...any) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	loweredQuery := strings.ToLower(query)

	if strings.Contains(loweredQuery, "clientes") {
		rows := sqlmock.NewRows([]string{"id", "nome", "limite", "saldo"})
		rows = rows.AddRow(1, "a", 0, 0)
		mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)
	} else if strings.Contains(loweredQuery, "transacoes") {
		rows := sqlmock.NewRows([]string{"id", "clienteid", "tipo", "valor", "descricao", "data_extrato"})
		rows = rows.AddRow(1, 1, "c", 500, "test", "00:00:00.000000")
		mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)
	}

	return db.Query(query)
}

func (dp *DatabaseProviderMock) Insert(query string, args ...any) error {
	return nil
}
