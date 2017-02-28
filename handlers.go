package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
    "github.com/TomyMMX/ad-server/models"
    "github.com/TomyMMX/ad-server/data"
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

	var ads []models.Ad
	var err error

	if vars["folderId"] == "" {
		//get all ads at the root of the folder structure
		ads, err = data.GetAds(0)
	} else {
		//here we are interested in the folder id
		folderId, _ := strconv.Atoi(vars["folderId"])

		//get all ads in the specified folder
		ads, err = data.GetAds(folderId)
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

	var folders []models.Folder
	var err error

	if vars["parrentId"] == "" {
		//get all folders at the root of the folder structure
		folders, err = data.GetFolders(0)
	} else {
		//here we are interested in the id of the parrent folder
		parrentId, _ := strconv.Atoi(vars["parrentId"])

		//if the Atoi function fails the parrentId will be 0
		//and we will return folders at root

		//get all folders in the specified parrent folder
		folders, err = data.GetFolders(parrentId)
	}

	PrepareAPIResponse(w, err)

	if err := json.NewEncoder(w).Encode(folders); err != nil {
		panic(err)
	}
}
func OneFolder(w http.ResponseWriter, r *http.Request) {}

func AddFolder(w http.ResponseWriter, r *http.Request) {

}
