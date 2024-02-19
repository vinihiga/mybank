package tests

import (
	"bytes"
	"encoding/json"
	statementController "mybank/internal/controllers/statement"
	"mybank/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStatement(t *testing.T) {
	// Given
	var buffer bytes.Buffer
	_ = json.NewEncoder(&buffer)

	var req = httptest.NewRequest(http.MethodGet, "/clientes/1/extrato", &buffer)
	var writer = httptest.NewRecorder()

	var sut = statementController.StatementController{}
	sut.DatabaseProvider = &mocks.DatabaseProviderMock{}

	// When
	sut.GetStatement(writer, req)

	res := writer.Result()
	defer res.Body.Close()

	// Then
	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
}
