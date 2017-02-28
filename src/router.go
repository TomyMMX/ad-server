package main

import ( 
    "net/http"

    /*Nice router from the gorilla web toolkit: http://www.gorillatoolkit.org/pkg/mux*/
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

    //use the mux router to find the correct handlers
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        //wrap the handler funcion with the logger
        handler = Logger(route.HandlerFunc, route.Name)    
    
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}