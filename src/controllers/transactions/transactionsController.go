package transactionsController

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	databaseProvider "mybank/src/providers/database"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const creditTransactionType string = "c"
const debtTransactionType string = "d"

type Transaction struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type Balance struct {
	id     int
	nome   string
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

// Adds the new transaction into dabase.
// PARAMETERS:
// w - The Responder Writer itself.
// r* - The pointer of the request where it'll be parsed as JSON.
func SetNewTransaction(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	var clientId string = mux.Vars(r)["id"]
	var newTransaction Transaction

	// First we need to verify if user can
	// transfer "credits" to outside of his/her
	// account.
	//
	// To do this, we need to decode the request
	// and check if it fits our business rules,
	// otherwise we should return an error code.
	if parseErr := json.NewDecoder(r.Body).Decode(&newTransaction); parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	result, notFoundUserErr := getBalance(clientId)
	var transaction string = strings.ToLower(newTransaction.Tipo)
	var isNewTransactionAboveLimit = newTransaction.Valor > (result.Saldo + result.Limite)

	if notFoundUserErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	} else if transaction == debtTransactionType && isNewTransactionAboveLimit {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte{})
		return
	}

	// Inserting the new transaction into the
	// database.
	insertErr := addNewTransaction(
		clientId,
		strings.ToLower(newTransaction.Tipo),
		newTransaction.Valor,
		newTransaction.Descricao,
	)

	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	// Now we need to get the data, because
	// we need to return the actual balance and
	// "limite" (credits).
	result, notFoundUserErr = getBalance(clientId)

	if notFoundUserErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	response, jsonError := json.Marshal(result)

	if jsonError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Adds new transaction into the database.
// PARAMETERS:
// clientId - The client's id
// transactionType - The values can be 'c' or 'd'. Where 'c' is credit and 'd' is debit.
// newValue - The value of the transaction.
// description - The description of the transaction.
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

	var insertErr error = databaseProvider.Shared.Insert(sql)

	if insertErr != nil {
		return insertErr
	}

	return nil
}

// Gets the current balance given the Client's ID.
// PARAMETERS:
// clientId - The client's id.
// RETURNS:
// *Balance - The balance's values, like limit and total.
// error - In case of the query failed or couldn't scan the data.
func getBalance(clientId string) (*Balance, error) {
	var sql string = fmt.Sprintf("SELECT * FROM clientes WHERE id = %s;", clientId)
	row := databaseProvider.Shared.Select(sql)

	var result Balance

	if row.Err() != nil {
		return nil, errors.New("couldn't find desired user")
	}

	if scanErr := row.Scan(&result.id, &result.nome, &result.Limite, &result.Saldo); scanErr != nil {
		return nil, errors.New("couldn't scan data")
	}

	return &result, nil
}
