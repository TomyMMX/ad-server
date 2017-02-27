package main

import ( 
    "fmt"
    "net/http" 
    "encoding/json"
    
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintln(w, "Ads REST API")
     //TODO: serve the API documentation or something
}

func AdsIndex(w http.ResponseWriter, r *http.Request) {
    //TODO: get the actual ads from the DB
    ads := Ads{
        Ad{Name: "A nice ad 1"},
        Ad{Name: "A nice ad 2"},
    }

    //since we know that we are returning JSON set the correct content type
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    //set the status code
    w.WriteHeader(http.StatusOK)
    
    if err := json.NewEncoder(w).Encode(ads); err != nil {
        panic(err)
    }
}

func OneAd(w http.ResponseWriter, r *http.Request) {
    //get the variables from the route
    vars := mux.Vars(r)
    //here we are interested in the ad id
    adId := vars["adId"]
    fmt.Fprintln(w, "Requested ad ID:", adId)
    
}