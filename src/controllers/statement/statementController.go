package statementController

import (
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
	Valor        int64
	Tipo         rune
	Descricao    string
	Realizada_em time.Time
}

type Balance struct {
	Total        int64
	Data_extrato time.Time
	Limite       int64
}

type Statement struct {
	Saldo              Balance
	Ultimas_transacoes []Transaction
}

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

func getBalance(clientId string) (*Balance, error) {
	var sql string = fmt.Sprintf("SELECT * FROM clientes WHERE id = %s;", clientId)
	row := databaseProvider.Select(sql)

	if row.Err() != nil {
		return nil, errors.New("Couldn't find specified client")
	}

	var balance Balance = Balance{
		Data_extrato: time.Now(),
	}

	row.Scan(&balance.Limite, &balance.Total)

	return &balance, nil
}

func getLastTransactions(clientId string) []Transaction {
	var sql string = fmt.Sprintf("SELECT * FROM transacoes WHERE clienteid = %s;", clientId)
	row, _ := databaseProvider.SelectMultiple(sql)

	var transactions []Transaction = make([]Transaction, 0)

	if row != nil {
		for row.Next() {
			var transaction = Transaction{}
			row.Scan(nil, nil, &transaction.Tipo, &transaction.Valor, nil)
			transactions = append(transactions, transaction)
		}
	}

	return transactions
}
