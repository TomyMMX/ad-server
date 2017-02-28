package data

import (
	//"database/sql"
	"errors"
    "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func PrepareDbConnection() (*sqlx.DB, error) {
	return sqlx.Open("mysql", "test:testpass@tcp(127.0.0.1:3306)/addb?parseTime=true")
}

func GetAds(folderId int) (Ads, error) {
	ads := Ads{}

	db, err := PrepareDbConnection()

	if err != nil {
		return ads, err
	}

	//get all ads in the desired folder
	err = db.Select(&ads, "SELECT * FROM ad WHERE folderid = ?", folderId)

	return ads, err
}

func GetFolders(parrentId int) (Folders, error) {
	folders := Folders{}

	db, err := PrepareDbConnection()

	if err != nil {
		return folders, err
	}

	//get all ads in the desired folder
	err = db.Select(&folders, "SELECT * FROM folder WHERE parrentid = ?", parrentId)

	return folders, err
}

func AddFolder(f Folder, parrentId int) error {
	db, err := PrepareDbConnection()

	if err != nil {
		return err
	}
    
    if(f.Name == ""){
        return errors.New("New folder name is empty.")
    }
    
    folderCount := 0
    err = db.Select(&folderCount, "SELECT count(*) FROM folder WHERE id = ?", parrentId)
    
    if folderCount == 0{
        return errors.New("Parrent folder with id " + strconv.Itoa(parrentId) + " does not exist.")
    }
    
    //also checked that this way of composing the sql query is safe against SQL injection
	//add this folder to the database
	_, err = db.Query("INSERT INTO folder (parrentid, name) VALUES (?, ?)", parrentId, f.Name)

	return err
}
