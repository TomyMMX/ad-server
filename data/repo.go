package data

import (
    //"database/sql"
    "errors"
    "strconv"

    "github.com/asaskevich/govalidator"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var connectionString string

func SetConnectionString(s string){
	connectionString = s
}

func PrepareDbConnection() (*sqlx.DB, error) {
    return sqlx.Open("mysql", connectionString)
}

/*HELPER FUNCTIONS*/
func CheckAd(a Ad) error {
    if a.Name == "" {
        return errors.New("Ad name is empty.")
    }
    
    if a.Url == "" {
        return errors.New("Ad URL is empty.")
    }
    
    if !govalidator.IsURL(a.Url) {
        return errors.New("Ad URL is invalid.")
    }
    
    return nil
}

/*DB ACCESS FUNCTIONS*/
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

func AddAd(a Ad) error {
    db, err := PrepareDbConnection()

    if err != nil {
        return err
    }
    
    if err=CheckAd(a); err != nil {
        return err;
    }
        
    var folderCount int
    err = db.Get(&folderCount, "SELECT count(*) FROM folder WHERE id = ?", a.FolderId)
    
    if folderCount == 0{
        return errors.New("Folder with id " + strconv.Itoa(a.FolderId) + " does not exist.")
    }
        
    //also checked that this way of composing the sql query is safe against SQL injection
    //add this folder to the database
    _, err = db.Query("INSERT INTO ad (folderid, name, url) VALUES (?, ?, ?)", a.FolderId, a.Name, a.Url)

    return err
}

func RemoveAd(adId int) error {
    db, err := PrepareDbConnection()

    if err != nil {
        return err
    }
    
    _, err = db.Query("DELETE FROM ad WHERE id=?", adId)

    return err
}

func UpdateAd(a Ad) error {
    db, err := PrepareDbConnection()

    if err != nil {
        return err
    }
    
    if err=CheckAd(a); err != nil {
        return err;
    }

    _, err = db.Query("UPDATE ad SET name=?, url=? WHERE id=?", a.Name, a.Url, a.Id)

    return err
}

func GetFolders(parentId int) (Folders, error) {
    folders := Folders{}

    db, err := PrepareDbConnection()

    if err != nil {
        return folders, err
    }

    //get all ads in the desired folder
    err = db.Select(&folders, "SELECT * FROM folder WHERE parentid = ?", parentId)

    return folders, err
}

func AddFolder(f Folder) error {
    db, err := PrepareDbConnection()

    if err != nil {
        return err
    }
    
    if f.Name == "" {
        return errors.New("New folder name is empty.")
    }
    
    //if adding int a existing folder, check if it exists
    if f.ParentId != 0 {
        folderCount := 0
        err = db.Get(&folderCount, "SELECT count(*) FROM folder WHERE id = ?", f.ParentId)
        
        if folderCount == 0{
            return errors.New("Parent folder with id " + strconv.Itoa(f.ParentId) + " does not exist.")
        }
    }
    
    //also checked that this way of composing the sql query is safe against SQL injection
    //add this folder to the database
    _, err = db.Query("INSERT INTO folder (parentid, name) VALUES (?, ?)", f.ParentId, f.Name)

    return err
}

func RemoveFolder(folderId int) error {
    db, err := PrepareDbConnection()

    if err != nil {
        return err
    }
    
    folderCount := 1
    err = db.Get(&folderCount, "SELECT count(*) FROM folder WHERE parentid = ?", folderId)
        
    if folderCount != 0{
        return errors.New("This folder contains at least one other folder. Delete that first.")
    }
    
    //begin transaction
    tx, err := db.Begin()
    
    //delete all ads that are in this folder
    _, err = tx.Exec("DELETE FROM ad WHERE folderId=?", folderId)
    if err != nil {
        tx.Rollback()
        return err;
    }
    
    //delete the folder
    _, err = tx.Exec("DELETE FROM folder WHERE id=?", folderId)
    if err != nil {
        tx.Rollback()
        return err;
    }
    
    //commit the whole transaction
    err = tx.Commit()

    return err
}

func UpdateFolder(f Folder) error {
    db, err := PrepareDbConnection()

    if err != nil {
        return err
    }
    
    if f.Name == "" {
        return errors.New("Folder name is empty.")
    }

    _, err = db.Query("UPDATE folder SET name=? WHERE id=?", f.Name, f.Id)

    return err
}