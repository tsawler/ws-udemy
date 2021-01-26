package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Staring web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
