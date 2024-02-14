package statementController

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	databaseProvider "mybank/src/providers/database"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Transaction struct {
	id           int
	clienteId    int
	Tipo         string    `json:"tipo"`
	Valor        int       `json:"valor"`
	Descricao    string    `json:"descricao"`
	Data_extrato time.Time `json:"realizada_em"`
}

type Balance struct {
	id           int
	nome         string
	Saldo        int       `json:"total"`
	Limite       int       `json:"limite"`
	Data_extrato time.Time `json:"data_extrato"`
}

type Statement struct {
	Saldo              Balance       `json:"saldo"`
	Ultimas_transacoes []Transaction `json:"ultimas_transacoes"`
}

// Gets the statement given a client id by passing into the url's query parameter.
// PARAMETERS:
// w - Response Writer.
// r* - The client's request that will be parsed into JSON.
func GetStatement(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	var clientId string = mux.Vars(r)["id"]
	result := Statement{}

	balance, notFoundUserErr := getBalance(clientId)
	transactions := getLastTransactions(clientId)

	if notFoundUserErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	result.Saldo = *balance
	result.Ultimas_transacoes = transactions

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

// Gets the current balance given the "clientId" (clienteId).
// PARAMETERS:
// - clientId: The client's ID.
// RETURNS:
// - *Balance: The result itself
// - error: In case of failed query.
func getBalance(clientId string) (*Balance, error) {
	var query string = fmt.Sprintf("SELECT * FROM clientes WHERE id = %s;", clientId)
	row := databaseProvider.Shared.Select(query)

	if row.Err() == sql.ErrNoRows {
		return nil, errors.New("couldn't find specified client")
	}

	var balance Balance = Balance{
		Data_extrato: time.Now(),
	}

	row.Scan(&balance.id, &balance.nome, &balance.Limite, &balance.Saldo)

	return &balance, nil
}

// Gets all transactions given the "clientId" (clienteId).
// PARAMETERS:
// - clientId: The client's ID.
// RETURNS:
// - []Transaction: Empty in case not found or with the respectively values.
func getLastTransactions(clientId string) []Transaction {
	var sql string = fmt.Sprintf("SELECT * FROM transacoes WHERE clienteid = %s;", clientId)
	rows, queryErr := databaseProvider.Shared.SelectMultiple(sql)
	var transactions []Transaction = make([]Transaction, 0)

	if queryErr != nil {
		return transactions
	}

	for rows.Next() {
		var transaction Transaction
		scanErr := rows.Scan(
			&transaction.id,
			&transaction.clienteId,
			&transaction.Tipo,
			&transaction.Valor,
			&transaction.Descricao,
			&transaction.Data_extrato,
		)

		if scanErr != nil {
			log.Printf(scanErr.Error())
			continue
		}

		transactions = append(transactions, transaction)
	}

	if rows.Err() != nil {
		return transactions
	}

	return transactions
}
