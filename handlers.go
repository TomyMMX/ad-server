package main

import ( 
    "fmt"
    "net/http"    
)

func Index(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintln(w, "Ads REST API")
     //TODO: serve the API documentation or something
}

func AdsIndex(w http.ResponseWriter, r *http.Request) {
    
}

func OneAd(w http.ResponseWriter, r *http.Request) {
    
}