package main

import (
	"net/http"

	"github.com/x022513319/asynchronous-processing-practice/api"
)

// API server entry point
func main() {

	r := api.NewRouter()

	http.ListenAndServe(":8080", r)
}
