package transactions

import (
	"encoding/json"
	"net/http"
)

type transaction struct {
	Valor     int64
	Tipo      rune
	Descricao string
}

type balance struct {
	Limite int64
	Saldo  int64
}

func SetNewTransaction(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	mock := balance{
		Limite: 1000,
		Saldo:  10000,
	}

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
