package data

import (
    "time"
	"errors"
	"strings"
	
	"github.com/asaskevich/govalidator"
)

type Folder struct {
    Id           int    `json:"id"`
    ParentId    int    `json:"parentid"` //if 0 then this folder is at the root level
    Name         string    `json:"name"`
    LastModified time.Time `json:"lastmodified"`
}

type Folders []Folder

func (f Folder) Check() (bool, error) {
	if f.Name == "" {
        return false, errors.New("Folder name is empty.")
    }
	
	if !IsValidName(f.Name) {
		return false, errors.New("Folder name is not a valid name.")
	}
    
    return true, nil
}

type Ad struct {
    Id           int    `json:"id"`
    FolderId     int    `json:"folderid"`
    Name         string    `json:"name"`
    Url          string    `json:"url"`
    LastModified time.Time `json:"lastmodified"`
}

type Ads []Ad

func (a Ad) Check() (bool, error) {
    if a.Name == "" {
        return false, errors.New("Ad name is empty.")
    }
	
	if !IsValidName(a.Name) {
		return false, errors.New("Ad name is not a valid name.")
	}
    
    if a.Url == "" {		
        return false, errors.New("Ad URL is empty.")
    }
    
    if !govalidator.IsURL(a.Url) {
        return false, errors.New("Ad URL is invalid.")
    }
    
    return true, nil
}

func IsValidName(s string)  bool {
	//TODO: check the string against a dictionary of invalid names and/or characters
	
	if strings.Contains(strings.ToLower(s), "badword"){
		return false
	}
	
	return true
}