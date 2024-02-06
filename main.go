package main

import (
	"log"
	"mybank/api/statement"
	"mybank/api/transactions"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/clientes/{id}/extrato", statement.GetStatement).Methods("GET")
	router.HandleFunc("/clientes/{id}/transacoes", transactions.SetNewTransaction).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
