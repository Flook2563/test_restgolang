package main

import (
	"log"

	"github.com/Flook2563/test_restgolang/db"
	"github.com/Flook2563/test_restgolang/models"
)

func main() {
	// Set Database
	db.InitDatabase()

	// Set Data to Struct
	customers := []models.Customer{
		{ID: 1, Name: "Nattanon Hanpap", Age: 26},
		{ID: 2, Name: "Natacha Goodperson", Age: 30},
		{ID: 3, Name: "Zreba Lastanime", Age: 50},
	}

	// Struct to database
	for _, customer := range customers {
		if err := db.GetDB().Create(&customer).Error; err != nil {
			log.Println("initially Error :", err)
		}
	}

	log.Println("Database was initially !!!")
}
