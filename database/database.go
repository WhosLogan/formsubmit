package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("formsubmit.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DbInstance = db
	return nil
}
