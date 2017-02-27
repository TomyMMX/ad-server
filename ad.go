package main

type Ad struct {
    Id        uint32    `json:"id"`  
    FolderId  uint32    `json:"folderid"`
    Name      string    `json:"name"`
    Url       string    `json:"url"`
    LastModified time.Time `json:"lastmodified"`
}

type Ads []Ad