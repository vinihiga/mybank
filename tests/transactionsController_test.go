package tests

import (
	"bytes"
	"encoding/json"
	transactionsController "mybank/src/controllers/transactions"
	"mybank/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

type transaction_mock struct {
	Valor     int
	Tipo      string
	Descricao string
}

func TestSetNewTransaction(t *testing.T) {
	// Given
	var mock = transaction_mock{
		Valor:     1000,
		Tipo:      "c",
		Descricao: "lorem ipsum",
	}

	var buffer bytes.Buffer
	_ = json.NewEncoder(&buffer).Encode(mock)

	var req = httptest.NewRequest(http.MethodPost, "/clientes/1/transacoes", &buffer)
	var writer = httptest.NewRecorder()

	var sut = transactionsController.TransactionsController{}
	sut.DatabaseProvider = &mocks.DatabaseProviderMock{}

	// When
	sut.SetNewTransaction(writer, req)

	res := writer.Result()
	defer res.Body.Close()

	// Then
}
