package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    /*Nice router from the gorilla web toolkit: http://www.gorillatoolkit.org/pkg/mux*/
    "github.com/gorilla/mux"
)

//basic server
func main() {
    
    //use the mux router to run the Index hanlder when the / endpoint is called
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    http.Handle("/", router)

    //run the server on por 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}