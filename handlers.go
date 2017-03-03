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

/*HELPER FUNCTIONS*/
func ReadRequestBody(r *http.Request) []byte {
    //limit reader so users can't flood us with large amounts of data
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    
    return body;
}

func RequestToAd (r *http.Request) (data.Ad, error) {
    var ad data.Ad
    
    body := ReadRequestBody(r)
    //unmarshal into our Ad struct
    err := json.Unmarshal(body, &ad)
    
    return ad, err
}

func RequestToFolder(r *http.Request) (data.Folder, error) {
    var folder data.Folder
   
    body := ReadRequestBody(r)
    //unmarshal into our Folder struct
    err := json.Unmarshal(body, &folder)
    
    return folder, err;
}

func PrepareAPIResponse(w http.ResponseWriter, err error, okStatus int) APIStatus{
    //since we know that we are returning JSON set the correct content type
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    
	var status APIStatus
    
    //set the status code
    if err != nil {
        //TODO: set correct response status for different error types
        w.WriteHeader(http.StatusBadRequest)
        status = APIStatus{
            Status: "Error",
            Reason: err.Error(),
        }    
		
		//in case of an error respond with the status JSON
        if err := json.NewEncoder(w).Encode(status); err != nil {
            panic(err)
        }
    } else {
        w.WriteHeader(okStatus)
        status = APIStatus{
            Status: "OK",            
        }    
    }
    
    return status
}

/*AD ENDPOINTS*/
func AdsInFolder(w http.ResponseWriter, r *http.Request) {
    //get the variables from the route
    vars := mux.Vars(r)

    //here we are interested in the folder id
    folderId, _ := strconv.Atoi(vars["folderId"])

    //get all ads in the specified folder
    ads, err := data.GetAds(folderId)
    
    if s := PrepareAPIResponse(w, err, http.StatusOK); s.Status == "OK" {
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

func AddAd(w http.ResponseWriter, r *http.Request) {
    ad, err := RequestToAd(r)
        
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
        
    vars := mux.Vars(r)
    ad.FolderId, err = strconv.Atoi(vars["folderId"])
    
    //parsing the folderId was not successful
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
        
    err = data.AddAd(ad)
    
    if s := PrepareAPIResponse(w, err, http.StatusCreated); s.Status == "OK" {
        s.Reason = "Successfully added new ad in folder: " + vars["folderId"]
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}

func RemoveAd(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)    
    adId, err := strconv.Atoi(vars["adId"])
    
    //parsing the adId was not successful
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    err = data.RemoveAd(adId)
    
    if s := PrepareAPIResponse(w, err, http.StatusOK); s.Status == "OK" {
        s.Reason = "Successfully removed ad with id: " + vars["adId"]
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}

func UpdateAd(w http.ResponseWriter, r *http.Request) {
    ad, err := RequestToAd(r)
    
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    vars := mux.Vars(r)    
    ad.Id, err = strconv.Atoi(vars["adId"])
    
    //parsing the folderId was not successful
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    err = data.UpdateAd(ad)
    
    if s := PrepareAPIResponse(w, err, http.StatusOK); s.Status == "OK" {
        s.Reason = "Successfully updated ad with id: " + vars["adId"]
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}

/*FOLDER ENDPOINTS*/
func FoldersInFolder(w http.ResponseWriter, r *http.Request) {
    //get the variables from the route
    vars := mux.Vars(r)

    //here we are interested in the id of the parent folder
    parentId, _ := strconv.Atoi(vars["parentId"])

    //if the Atoi function fails the parentId will be 0
    //and we will return folders at root

    //get all folders in the specified parent folder
    folders, err := data.GetFolders(parentId)
    
    if s := PrepareAPIResponse(w, err, http.StatusOK); s.Status == "OK" {
        if err := json.NewEncoder(w).Encode(folders); err != nil {
            panic(err)
        }
    }
}
func OneFolder(w http.ResponseWriter, r *http.Request) {
    //get the variables from the route
    vars := mux.Vars(r)
    //here we are interested in the ad id
    folderId := vars["folderId"]
    fmt.Fprintln(w, "Requested folder ID:", folderId)
    
    //TODO: implement return of one specific folder
}

func AddFolder(w http.ResponseWriter, r *http.Request) {
    folder, err := RequestToFolder(r)
        
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    vars := mux.Vars(r)       
    if vars["parentId"] == "" {
        folder.ParentId = 0
    } else {
        folder.ParentId, err = strconv.Atoi(vars["parentId"])
    }
    
    //parsing the parentId was not successful
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
        
    err = data.AddFolder(folder)
    
    if s := PrepareAPIResponse(w, err, http.StatusCreated); s.Status == "OK" {
        s.Reason = "Successfully added new folder in parent: " + vars["parentId"]
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}

func RemoveFolder(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)    
    folderId, err := strconv.Atoi(vars["folderId"])
    
    //parsing the folderId was not successful
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    err = data.RemoveFolder(folderId)
    
    if s := PrepareAPIResponse(w, err, http.StatusOK); s.Status == "OK" {
        s.Reason = "Successfully removed folder with id: " + vars["folderId"]
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}

func UpdateFolder(w http.ResponseWriter, r *http.Request) {
    folder, err := RequestToFolder(r)
    
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    vars := mux.Vars(r)    
    folder.Id, err = strconv.Atoi(vars["folderId"])
    
    //parsing the folderId was not successful
    if err != nil {
        PrepareAPIResponse(w, err, 0)
        return
    }
    
    err = data.UpdateFolder(folder)
    
    if s := PrepareAPIResponse(w, err, http.StatusOK); s.Status == "OK" {
        s.Reason = "Successfully updated folder with id: " + vars["folderId"]
        if err := json.NewEncoder(w).Encode(s); err != nil {
            panic(err)
        }
    }
}