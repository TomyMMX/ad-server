package main

import (
    "log"
    "net/http"
)

//basic web server that handles API requests on our ad server
func main() {    
    router := NewRouter()

    //run the server on por 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}