package main

import (
	"log"
	"net/http"
)

//some configs
//const

//basic web server that handles API requests on our ad server
func main() {
	router := NewRouter()

	//run the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
