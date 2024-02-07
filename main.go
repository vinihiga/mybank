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

	var port string = ":27000"

	router := mux.NewRouter()
	router.HandleFunc("/clientes/{id}/extrato", statement.GetStatement).Methods("GET")
	router.HandleFunc("/clientes/{id}/transacoes", transactions.SetNewTransaction).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Default().Printf("Server starting at internal port %s!\n", server.Addr)

	log.Fatal(server.ListenAndServe())

	log.Default().Printf("Server started!\n")
}
