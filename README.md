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

## API Reference
**Working with ad folders**  
*GET /api/folders*  
Get a list of all folders at root level.

*GET /api/folders/parent/{parentId}*  
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

*POST /api/folders*  
Add a folder at root level.

*POST /api/folders/parent/{parentId}*  
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

*PUT /api/folders/{folderId}*  
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

*DELETE /api/folders/{folderId}*  
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
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 01 Mar 2017 13:11:03 GMT
Content-Length: 107
{
"status": "Error",
"reason": "This folder contains at least one other folder. Delete that first."
}
```

**Working with ads**  
