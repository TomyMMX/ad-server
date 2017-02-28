package main

type Folder struct {
	Id           uint32    `json:"id"`
	ParentId     uint32    `json:"parentid"` //if 0 then this folder is at the root level
	Name         string    `json:"name"`
	LastModified time.Time `json:"lastmodified"`
}

type Folders []Folder
