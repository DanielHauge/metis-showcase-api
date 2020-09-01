package main

import (
	cors "github.com/rs/cors"
	"log"
	"net/http"
)



func main() {

	log.Println("Initializing API.")
	router := NewRouter()
	handler := cors.Default().Handler(router)


	log.Fatal(http.ListenAndServe(":8787", handler))
	log.Println("Server running on port 8787.")
}
