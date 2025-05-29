package main

import (
	"log"
	"net/http"
	"portal/routes"
)

func main() {
	router := routes.SetupRouter()

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
