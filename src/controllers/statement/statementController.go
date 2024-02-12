package statementController

import (
	"encoding/json"
	"log"
	databaseProvider "mybank/src/providers/database"
	"net/http"
	"time"
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

type test struct {
	Id    int
	Nome  string
	Saldo int
}

func GetStatement(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	//vars := mux.Vars(r)
	mock := Statement{
		Saldo: Balance{
			Total:        1,
			Data_extrato: time.Now(),
			Limite:       1,
		},
		Ultimas_transacoes: make([]Transaction, 0),
	}

	row := databaseProvider.Select("SELECT * FROM clientes;")

	if row.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	var t test
	row.Scan(&t.Id, &t.Nome, &t.Saldo)

	mock.Saldo.Limite = int64(t.Saldo)
	response, error := json.Marshal(mock)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
