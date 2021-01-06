package main

import (
	"net/http"

	"go-http-testing/app"
)

func main() {

	http.ListenAndServe(":8000", app.NewHandler())
}
