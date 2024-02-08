package statementController

import (
	"encoding/json"
	"log"
	databaseProvider "mybank/src/providers/database"
	"net/http"
	"time"
)

type transaction struct {
	Valor        int64
	Tipo         rune
	Descricao    string
	Realizada_em time.Time
}

type balance struct {
	Total        int64
	Data_extrato time.Time
	Limite       int64
}

type statement struct {
	Saldo              balance
	Ultimas_transacoes []transaction
}

type test struct {
	Id    int
	Nome  string
	Saldo int
}

func GetStatement(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("Received request")

	//vars := mux.Vars(r)
	mock := statement{
		Saldo: balance{
			Total:        1,
			Data_extrato: time.Now(),
			Limite:       1,
		},
		Ultimas_transacoes: make([]transaction, 0),
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
