package transactionsController

import (
	"encoding/json"
	"fmt"
	"log"
	databaseProvider "mybank/src/providers/database"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction struct {
	Valor     int64  // Value in english
	Tipo      string // Type in english (can be "C" for credit or "D" for debit)
	Descricao string // Description in english
}

type Balance struct {
	Limite int64 // Account "extra credit"
	Saldo  int64 // Balance itself
}

func SetNewTransaction(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	var clientId string = mux.Vars(r)["id"]
	var newTransaction Transaction
	var result Balance

	parseErr := json.NewDecoder(r.Body).Decode(&newTransaction)

	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	var sql = fmt.Sprintf(
		"INSERT INTO transacoes (clienteid, tipo, valor) VALUES (%s, '%s', %d);",
		clientId,
		newTransaction.Tipo,
		newTransaction.Valor,
	)

	var insertErr error = databaseProvider.Insert(sql)

	if insertErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	sql = fmt.Sprintf("SELECT * FROM clientes WHERE id = %s;", clientId)
	row := databaseProvider.Select(sql)

	if row.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	row.Scan(&result.Limite, &result.Saldo)
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
