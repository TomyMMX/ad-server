#Ad REST API
##Description
REST API written in Go that provides different types of clients the abillity to perform CRUD operations on ads.  
Ads in this case are simple (named) URLs organized into folders. Each folder can contain an arbitrary number of ads and subfolders.

##Technology stack
Written in Go.  
The main reason beeing the simple nature of the task and Go beeing really good ad doing simple things really fast.  
For most things the standard libraries are used the few exceptions are:
* github.com/gorilla/mux - Nice router from the gorilla web toolkit: http://www.gorillatoolkit.org/pkg/mux 
* github.com/go-sql-driver/mysql - A mysql driver for Go
* github.com/jmoiron/sqlx - provides a set of extensions on go's standard database/sql library. Mainly usefull because it provides marshalling of rows into structs.
* github.com/asaskevich/govalidator - for validating the ad URLs

The database used is MySQL. No special reason for that... could have been any relational database. Actually changing the database would not require much work.

##Project organization  
The main package is divided into two .go files. main.go and hanlders.go. This code provides tha backbone of the API server with all the necesery routes and hanlder functions to handle specific routes.  
Other parts of the application are divided into 3 additional packages where the data package is the most important one. It contains code for our data models and code that communicates with the database.  
The database schema for the rather simple two table database can be found in [DB/dbschema.sql](https://github.com/TomyMMX/ad-server/blob/master/DB/dbschema.sql).  

###Testing
The only tests I could think of are at the model/data layer. Where tests check if the functions that check the validity of our Ad or Folder objects work for different edge type inputs.  
The other thing here that should be tested are the functions that work with the database. So basicaly to check if adding, retreaving, updating and deleting works.  
My proficiency in Go is at this point a bit limited so I have no idea how to test the handler functions that work on the http request and respond to it. 

## API Reference

**Working with ad folders**  

**GET** */api/folders*  
Get a list of all folders at root level.

**GET** */api/folders/parent/{parentId}*  
Get a list of all folders inside the folder with the id {parentId}.

**Request:**
```JSON
GET /api/folders/parent/1 HTTP/1.1
Accept: application/json
```
**Response:**
```JSON
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 12:43:32 GMT
Content-Length: 79
[
{
  "id":2,
  "parentid":1,
  "name":"Subfolder",
  "lastmodified":"2017-03-01T00:39:08Z"
}
]
```

**POST** */api/folders*  
Add a folder at root level.

**POST** */api/folders/parent/{parentId}*  
Add a folder inside the fodler with id {parentId}.

**Request:**
```JSON
POST /api/folders/parent/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
{
  "name": "New folder"
}
```

**Response:**
```JSON
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:04:11 GMT
Content-Length: 81
{
  "status": "OK",
  "reason": "Successfully added new folder in parent: 1"
}
```

**PUT** */api/folders/{folderId}*  
Update the folder with id {folderId}.

**Request:**
```JSON
PUT /api/folders/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
{
  "name": "Updated folder name"
}
```

**Response:**
```JSON
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:08:30 GMT
Content-Length: 77
{
  "status": "OK",
  "reason": "Successfully updated folder with id: 1"
}
```

**DELETE** */api/folders/{folderId}*  
Remove the folder with id {folderId}.

**Request:**
```JSON
DELETE /api/folders/1 HTTP/1.1
Accept: application/json
```

**Response:**
```JSON
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:08:30 GMT
Content-Length: 77
{
  "status": "OK",
  "reason": "Successfully removed folder with id: 1"
}
```
If folder has at least one subfolder the DELETE action will fail.
```JSON
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:11:03 GMT
Content-Length: 107
{
  "status": "Error",
  "reason": "This folder contains at least one other folder. Delete that first."
}
```

**Working with ads**  

**GET** */api/ads/folder/{folderId}*  
Get a list of all ads inside the folder with the id {folderId}.

**Request:**
```JSON
GET /api/ads/folder/1 HTTP/1.1
Accept: application/json
```
**Response:**
```JSON
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:17:27 GMT
Content-Length: 145
[
  {
  "id": 10,
  "folderid": 1,
  "name": "Ace of hearts",
  "url": "http://cliparts.co/cliparts/8c6/oga/8c6ogaBoi.png",
  "lastmodified": "2017-03-01T01:42:50Z"
  }
]
```

**POST** */api/ads/folder/{folderId}*  
Add a ad inside the fodler with id {folderId}.

**Request:**
```JSON
POST /api/ads/folder/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
{
  "name": "New cool ad",
  "url": "http://www.hdbloggers.net/wp-content/uploads/2016/04/Cool-Wallpapers.jpg"
}
```
URLs get tested for validity. URLs that are not URLs will not be accepted.

**Response:**
```JSON
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:20:45 GMT
Content-Length: 77
{
  "status": "OK",
  "reason": "Successfully added new ad in folder: 1"
}
```

**PUT** */api/ads/{adId}*  
Update the ad with id {adId}.

**Request:**
```JSON
PUT /api/ads/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
{
  "name": "New cool ad - UPDATE",
  "url": "this is not an URL"
}
```

**Response:**
```JSON
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:25:17 GMT
Content-Length: 59
{
  "status": "Error",
  "reason": "Ad URL is invalid."
}
```

**DELETE** */api/ads/{adId}*  
Remove the ad with id {adId}.

**Request:**
```JSON
DELETE /api/ads/1 HTTP/1.1
Accept: application/json
```

**Response:**
```JSON
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:26:44 GMT
Content-Length: 73
{
  "status": "OK"
  "reason": "Successfully removed ad with id: 1"
}
```
