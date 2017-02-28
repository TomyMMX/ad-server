package main

import (
	"log"
	"net/http"

	"github.com/TomyMMX/ad-server/router"
)

//the routes for this API server
var routes = router.Routes{
	router.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	router.Route{
		"AdsIndex",
		"GET",
		"/api/ads",
		AdsInFolder,
	},
	router.Route{
		"AdsInFolder",
		"GET",
		"/api/ads/folder/{folderId}",
		AdsInFolder,
	},
    router.Route{
		"AddAdInFolder",
		"POST",
		"/api/ads/folder/{folderId}",
		AddAd,
	},
    router.Route{
		"RemoveAd",
		"DELETE",
		"/api/ads/{adId}",
		RemoveAd,
	},
	router.Route{
		"OneAd",
		"GET",
		"/api/ads/{adId}",
		OneAd,
	},
	router.Route{
		"FoldersIndex",
		"GET",
		"/api/folders",
		FoldersInFolder,
	},
	router.Route{
		"AddFolder",
		"POST",
		"/api/folders",
		AddFolder,
	},
	router.Route{
		"FoldersInFolder",
		"GET",
		"/api/folders/parrent/{parrentId}",
		FoldersInFolder,
	},
	router.Route{
		"AddFolderInFolder",
		"POST",
		"/api/folders/parrent/{parrentId}",
		AddFolder,
	},
	router.Route{
		"OneFodler",
		"GET",
		"/api/folders/{folderId}",
		OneFolder,
	},
}

//basic web server that handles API requests on our ad server
func main() {

	r := router.NewRouter(routes)

	//run the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}
