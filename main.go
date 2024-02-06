package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/balance", getBalance)

	http.ListenAndServe(":80", nil)
}
