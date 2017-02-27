package main

import (
    "net/http"    
)

//the structure of our routes
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

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
        "/api/ads",
        AdsIndex,
    },
    Route{
        "OneAd",
        "GET",
        "/api/ads/{adId}",
        OneAd,
    },
}