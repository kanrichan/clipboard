package model

import (
	"os"
	"path"

	_ "github.com/fumiama/sqlite3"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Setup() error {
	const dbpath = "./data/database.db"

	var p = path.Dir(dbpath)
	if _, err := os.Stat(p); err != nil || os.IsNotExist(err) {
		err := os.Mkdir(p, 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(dbpath); err != nil || os.IsNotExist(err) {

		f, err := os.Create(dbpath)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	db, err := gorm.Open("sqlite3", dbpath)
	if err != nil {
		return err
	}

	db.AutoMigrate(&User{})

	DB = db
	return nil
}
