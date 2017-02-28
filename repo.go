package main

import (
	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func PrepareDbConnection() (*sqlx.DB, error) {
	return sqlx.Open("mysql", "test:testpass@tcp(127.0.0.1:3306)/addb?parseTime=true")
}

func GetAds(folderId int) ([]Ad, error) {
	ads := []Ad{}

	db, err := PrepareDbConnection()

	if err != nil {
		return ads, err
	}

	//get all ads in the desired folder
	err = db.Select(&ads, "SELECT * FROM ad WHERE folderid = ?", folderId)

	return ads, err
}
