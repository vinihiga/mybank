package api

import (
	"fmt"
	"net/http"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1234.00")
}
