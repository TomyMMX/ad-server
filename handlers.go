package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "io"
    "io/ioutil"

	"github.com/gorilla/mux"
    "github.com/TomyMMX/ad-server/data"
)

type APIStatus struct {
	Status  string  `json:"status"`
    Code    string  `json:"code"`
    Reason  string  `json:"reason"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ads REST API")
	//TODO: serve the API documentation or something
}

func PrepareAPIResponse(w http.ResponseWriter, err error) APIStatus{
	//since we know that we are returning JSON set the correct content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    var status APIStatus
    
	//set the status code
	if err != nil {
        //TODO: set correct response status for different error types
		w.WriteHeader(http.StatusBadRequest)
        status = APIStatus{
            Status: "Error",
            Reason: err.Error(),
        }    

        if err := json.NewEncoder(w).Encode(status); err != nil {
            panic(err)
        }
	} else {
		w.WriteHeader(http.StatusOK)
        status = APIStatus{
            Status: "OK",            
        }    
	}
    
    return status
}

func AdsInFolder(w http.ResponseWriter, r *http.Request) {
	//get the variables from the route
	vars := mux.Vars(r)

	var ads []data.Ad
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

	if s := PrepareAPIResponse(w, err); s.Status == "OK" {
        if err := json.NewEncoder(w).Encode(ads); err != nil {
            panic(err)
        }
    }
}

func OneAd(w http.ResponseWriter, r *http.Request) {
	//get the variables from the route
	vars := mux.Vars(r)
	//here we are interested in the ad id
	adId := vars["adId"]
	fmt.Fprintln(w, "Requested ad ID:", adId)
    
    //TODO: implement return of one specific ad
}

func FoldersInFolder(w http.ResponseWriter, r *http.Request) {
	//get the variables from the route
	vars := mux.Vars(r)

	var folders []data.Folder
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

	if s := PrepareAPIResponse(w, err); s.Status == "OK" {
        if err := json.NewEncoder(w).Encode(folders); err != nil {
            panic(err)
        }
    }
}
func OneFolder(w http.ResponseWriter, r *http.Request) {}

func ReadRequestBody(r *http.Request) []byte {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    
    return body;
}

func AddFolder(w http.ResponseWriter, r *http.Request) {
    var folder data.Folder
    
    //get the variables from the route
	vars := mux.Vars(r)
    
    body := ReadRequestBody(r)
    //unmarshal into our Folder struct
    err := json.Unmarshal(body, &folder)
        
    if err != nil {
        PrepareAPIResponse(w, err)
        return
    }
    
    var parrentId int
    if vars["parrentId"] == "" {
        parrentId = 0
    } else {
        parrentId, err = strconv.Atoi(vars["parrentId"])
    }
    
    //parsing the parrentId was not successful
    if err != nil {
        PrepareAPIResponse(w, err)
        return
    }
        
    err = data.AddFolder(folder, parrentId)
    
    if s := PrepareAPIResponse(w, err); s.Status == "OK" {
        s.Reason = "Successfully added new folder in parrent: "+strconv.Itoa(parrentId)
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}
