package data

import (
    "time"
)

type Folder struct {
    Id           int    `json:"id"`
    ParrentId    int    `json:"parrentid"` //if 0 then this folder is at the root level
    Name         string    `json:"name"`
    LastModified time.Time `json:"lastmodified"`
}

type Folders []Folder

type Ad struct {
    Id           int    `json:"id"`
    FolderId     int    `json:"folderid"`
    Name         string    `json:"name"`
    Url          string    `json:"url"`
    LastModified time.Time `json:"lastmodified"`
}

type Ads []Ad