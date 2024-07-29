package main

import (
	"apiRest/app"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", app.Router()))
}
