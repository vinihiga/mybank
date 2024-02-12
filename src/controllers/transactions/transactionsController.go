package transactionsController

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	databaseProvider "mybank/src/providers/database"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction struct {
	Valor     int    // Value in english
	Tipo      string // Type in english (can be "C" for credit or "D" for debit)
	Descricao string // Description in english
}

type Balance struct {
	Limite int // Account "extra credit"
	Saldo  int // Balance itself
}

func SetNewTransaction(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	var clientId string = mux.Vars(r)["id"]
	var newTransaction Transaction

	parseErr := json.NewDecoder(r.Body).Decode(&newTransaction)

	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	insertErr := addNewTransaction(
		clientId,
		newTransaction.Tipo,
		newTransaction.Valor,
		newTransaction.Descricao,
	)

	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	result, notFoundUserErr := getBalance(clientId)

	if notFoundUserErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	response, error := json.Marshal(result)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func addNewTransaction(
	clientId string,
	transactionType string,
	newValue int,
	description string,
) error {

	var sql = fmt.Sprintf(
		"INSERT INTO transacoes (clienteid, tipo, valor, descricao) VALUES (%s, '%s', %d, '%s');",
		clientId,
		transactionType,
		newValue,
		description,
	)

	var insertErr error = databaseProvider.Insert(sql)

	if insertErr != nil {
		return insertErr
	}

	return nil
}

func getBalance(clientId string) (*Balance, error) {
	var sql string = fmt.Sprintf("SELECT * FROM clientes WHERE id = %s;", clientId)
	row := databaseProvider.Select(sql)

	var result Balance

	if row.Err() != nil {
		return nil, errors.New("couldn't find desired user")
	}

	row.Scan(&result.Limite, &result.Saldo)

	return &result, nil
}
