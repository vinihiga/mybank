package main

import (
	"fmt"
	"net/http"
)

func getBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1234.00")
}
