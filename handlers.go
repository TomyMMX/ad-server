package main

import ( 
    "html"
    "fmt"
    "net/http"    
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func AdsIndex(w http.ResponseWriter, r *http.Request) {
    
}

func OneAd(w http.ResponseWriter, r *http.Request) {
    
}