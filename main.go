package main

import (
	api "mybank/api"
	"net/http"
)

func main() {

	http.HandleFunc("/balance", api.GetBalance)

	http.ListenAndServe(":80", nil)
}
