package statement

import (
	"encoding/json"
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

func GetStatement(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	mock := statement{
		Saldo: balance{
			Total:        1,
			Data_extrato: time.Now(),
			Limite:       1,
		},
		Ultimas_transacoes: make([]transaction, 0),
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
