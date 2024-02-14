package main

import (
	"log"
	statementController "mybank/src/controllers/statement"
	transactionsController "mybank/src/controllers/transactions"
	databaseProvider "mybank/src/providers/database"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// When we start, we must setup the database in order
	// to use local instance or the cluster's one.
	if slices.Contains(os.Args, "--dev") {
		databaseProvider.Shared.SetupLocalEnvironment()
	} else {
		databaseProvider.Shared.SetupNormalEnvironment()
	}

	// Setting-up endpoints and its respectively controllers.
	// By default every endpoint will use port `27000` as decided below.
	var port string = ":27000"

	router := mux.NewRouter()
	router.HandleFunc("/clientes/{id}/extrato", statementController.GetStatement).Methods("GET")
	router.HandleFunc("/clientes/{id}/transacoes", transactionsController.SetNewTransaction).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	// Booting the server.
	// We are logging the port to help while we instantiate
	// with docker-compose or clusterization technique.
	log.Default().Printf("Server starting at internal port %s!\n", server.Addr)
	log.Fatal(server.ListenAndServe())
	log.Default().Printf("Server started!\n")
}
