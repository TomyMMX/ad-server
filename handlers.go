package main

import (
	"encoding/json"
	"fmt"
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

func AdsInFolder(w http.ResponseWriter, r *http.Request) {
	//get the variables from the route
	vars := mux.Vars(r)

	var ads []Ad
	var err error

	if vars["folderId"] == "" {
		//get all ads at the root of the folder structure
		ads, err = GetAds(0)
	} else {
		//here we are interested in the folder id
		folderId, _ := strconv.Atoi(vars["folderId"])

		//get all ads in the specified folder
		ads, err = GetAds(folderId)
	}

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

func FoldersInFolder(w http.ResponseWriter, r *http.Request) {
	//get the variables from the route
	vars := mux.Vars(r)

	var folders []Folder
	var err error

	if vars["parrentId"] == "" {
		//get all folders at the root of the folder structure
		folders, err = GetFolders(0)
	} else {
		//here we are interested in the id of the parrent folder
		parrentId, _ := strconv.Atoi(vars["parrentId"])

		//get all folders in the specified parrent folder
		folders, err = GetFolders(parrentId)
	}

	PrepareAPIResponse(w, err)

	if err := json.NewEncoder(w).Encode(folders); err != nil {
		panic(err)
	}
}
func OneFolder(w http.ResponseWriter, r *http.Request) {}
