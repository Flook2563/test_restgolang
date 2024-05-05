package main

import (
	"github.com/Flook2563/test_restgolang/db"
	"github.com/Flook2563/test_restgolang/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	
	db.InitDatabase()

	router := gin.Default()

	router.POST("/customers", handlers.CreateCustomer)
	router.GET("/customers", handlers.GetAllCustomers)
	router.GET("/customers/:id", handlers.GetCustomer)
	router.PUT("/customers/:id", handlers.UpdateCustomer)
	router.DELETE("/customers/:id", handlers.DeleteCustomer)

	if err := router.Run(":8080"); err != nil {
		panic("Fail run server : " + err.Error())
	}
}
