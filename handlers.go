package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ads REST API")
	//TODO: serve the API documentation or something
}

func PrepareAPIResponse(w http.ResponseWriter, err error) {
	//since we know that we are returning JSON set the correct content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//set the status code
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func AdsIndex(w http.ResponseWriter, r *http.Request) {
	//get all ads at the root of the folder structure
	ads, err := GetAds(0)

	PrepareAPIResponse(w, err)

	if err := json.NewEncoder(w).Encode(ads); err != nil {
		panic(err)
	}
}

func AdsInFolder(w http.ResponseWriter, r *http.Request) {
	//get the variables from the route
	vars := mux.Vars(r)
	//here we are interested in the folder id
	folderIds := vars["folderId"]

	log.Printf(
		"%s\t%s\t%s\t%s%s",
		r.Method,
		r.RequestURI,
		"AdsInFolder",
		"Requested ads of folder:",
		folderIds,
	)

	folderId, _ := strconv.Atoi(folderIds)

	//get all ads in the specified folder
	ads, err := GetAds(folderId)

	PrepareAPIResponse(w, err)

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
