package db

import (
	"log"

	"github.com/Flook2563/test_restgolang/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("customers.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect to the database Fail :", err)
	} else {
		log.Println("Connect to the database . . .")
	}

	dberr := db.AutoMigrate(&models.Customer{})
	if dberr != nil {
		log.Println("Auto-Migrate Fail :", dberr.Error())
	} else {
		log.Println("Auto-Migrate . . .")
	}

}

func GetDB() *gorm.DB {
	return db
}
