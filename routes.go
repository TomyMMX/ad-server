package main

import (
	"net/http"
)

//the structure of our routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AdsIndex",
		"GET",
		"/api/ads",
		AdsInFolder,
	},
	Route{
		"AdsInFolder",
		"GET",
		"/api/ads/folder/{folderId}",
		AdsInFolder,
	},
	Route{
		"OneAd",
		"GET",
		"/api/ads/{adId}",
		OneAd,
	},
	Route{
		"FoldersIndex",
		"GET",
		"/api/folders",
		FoldersInFolder,
	},
	Route{
		"FoldersInFolder",
		"GET",
		"/api/folders/parrent/{parrentId}",
		FoldersInFolder,
	},
	Route{
		"OneFodler",
		"GET",
		"/api/folders/{fodlerId}",
		OneFolder,
	},
}
