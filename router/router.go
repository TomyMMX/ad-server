package router

import (
    "net/http"

    /*Nice router from the gorilla web toolkit: http://www.gorillatoolkit.org/pkg/mux*/
    "github.com/gorilla/mux"    
    "github.com/TomyMMX/ad-server/logger"
)

//the structure of our routes
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {

    //use the mux router to find the correct handlers
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        //wrap the handler funcion with the logger
        handler = logger.Logger(route.HandlerFunc, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
