// main.go

package main

import (
	"log"
	"net/http"
	"crud-api/router"
)

func main() {
	r := router.InitializeRouter()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
