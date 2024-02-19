package main

import (
	"log"
	statementController "mybank/internal/controllers/statement"
	transactionsController "mybank/internal/controllers/transactions"
	databaseProvider "mybank/internal/providers/database"
	"net/http"
	"os"
	"slices"
)

func main() {

	// When we start, we must setup the database in order
	// to use local instance or the cluster's one.
	var databaseProvider databaseProvider.IDatabaseProvider = &databaseProvider.DatabaseProvider{}

	if slices.Contains(os.Args, "--dev") {
		databaseProvider.SetupLocalEnvironment()
	} else {
		databaseProvider.SetupNormalEnvironment()
	}

	// Setting-up endpoints and its respectively controllers.
	// By default every endpoint will use port `27000` as decided below.
	var addr string = "localhost:27000"

	mux := http.NewServeMux()

	var transactionsController transactionsController.TransactionsController
	transactionsController.DatabaseProvider = databaseProvider

	var statementController statementController.StatementController
	statementController.DatabaseProvider = databaseProvider

	mux.HandleFunc("POST /clientes/{id}/transacoes", transactionsController.SetNewTransaction)
	mux.HandleFunc("GET /clientes/{id}/extrato", statementController.GetStatement)

	// Booting the server.
	// We are logging the port to help while we instantiate
	// with docker-compose or clusterization technique.
	log.Default().Printf("Server starting at %s!\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
	log.Default().Printf("Server started!\n")
}
