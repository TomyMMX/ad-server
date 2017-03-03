package data

import (
    "time"
	"errors"
	
	"github.com/asaskevich/govalidator"
)

type Folder struct {
    Id           int    `json:"id"`
    ParentId    int    `json:"parentid"` //if 0 then this folder is at the root level
    Name         string    `json:"name"`
    LastModified time.Time `json:"lastmodified"`
}

type Folders []Folder

func (f Folder) Check() error {
	if f.Name == "" {
        return errors.New("Folder name is empty.")
    }
	
	if !IsValidName(f.Name) {
		return errors.New("Folder name is not a valid name.")
	}
    
    return nil
}

type Ad struct {
    Id           int    `json:"id"`
    FolderId     int    `json:"folderid"`
    Name         string    `json:"name"`
    Url          string    `json:"url"`
    LastModified time.Time `json:"lastmodified"`
}

type Ads []Ad

func (a Ad) Check() error {
    if a.Name == "" {
        return errors.New("Ad name is empty.")
    }
	
	if !IsValidName(a.Name) {
		return errors.New("Ad name is not a valid name.")
	}
    
    if a.Url == "" {
        return errors.New("Ad URL is empty.")
    }
    
    if !govalidator.IsURL(a.Url) {
        return errors.New("Ad URL is invalid.")
    }
    
    return nil
}

func IsValidName(s string)  bool {
	//TODO: check the string against a dictionary of invalid names and/or characters
	return true
}