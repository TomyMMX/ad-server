package main

import (
	//"database/sql"
	//"log"

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
