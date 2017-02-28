package main

import (
	"time"
)

type Folder struct {
	Id           uint32    `json:"id"`
	ParrentId    uint32    `json:"parrentid"` //if 0 then this folder is at the root level
	Name         string    `json:"name"`
	LastModified time.Time `json:"lastmodified"`
}

type Folders []Folder
