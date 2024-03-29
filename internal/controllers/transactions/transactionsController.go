package transactionsController

import (
	"encoding/json"
	"errors"
	"log"
	databaseProvider "mybank/internal/providers/database"
	"net/http"
	"strings"
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

type TransactionsController struct {
	DatabaseProvider databaseProvider.IDatabaseProvider
}

// Adds the new transaction into dabase.
// PARAMETERS:
// w - The Responder Writer itself.
// r* - The pointer of the request where it'll be parsed as JSON.
func (controller *TransactionsController) SetNewTransaction(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	var path string = r.URL.Path
	var segments []string = strings.Split(path, "/")
	var clientId string = segments[2]
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

	result, notFoundUserErr := controller.getBalance(clientId)
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
	insertErr := controller.addNewTransaction(
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
	result, notFoundUserErr = controller.getBalance(clientId)

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
func (controller *TransactionsController) addNewTransaction(
	clientId string,
	transactionType string,
	newValue int,
	description string,
) error {

	var insertErr error = controller.DatabaseProvider.Insert(
		"INSERT INTO transacoes (clienteid, tipo, valor, descricao) VALUES ($1, $2, $3, $4)",
		clientId,
		transactionType,
		newValue,
		description,
	)

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
func (controller *TransactionsController) getBalance(clientId string) (*Balance, error) {
	rows, queryErr := controller.DatabaseProvider.Select("SELECT * FROM clientes WHERE id = $1", clientId)

	var result Balance

	if queryErr != nil {
		return nil, errors.New("couldn't find desired user")
	} else if !rows.Next() {
		return nil, errors.New("couldn't find desired user")
	} else if scanErr := rows.Scan(&result.id, &result.nome, &result.Limite, &result.Saldo); scanErr != nil {
		return nil, errors.New("couldn't scan data")
	}

	return &result, nil
}
