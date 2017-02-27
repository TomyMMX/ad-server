package main

import (
    "net/http"    
    /*Nice router from the gorilla web toolkit: http://www.gorillatoolkit.org/pkg/mux*/
    "github.com/gorilla/mux"
)

//the structure of our routes
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {

    //use the mux router to find the correct handlers
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    return router
}

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "AdsIndex",
        "GET",
        "/ads",
        AdsIndex,
    },
    Route{
        "OneAd",
        "GET",
        "/ads/{adId}",
        OneAd,
    },
}