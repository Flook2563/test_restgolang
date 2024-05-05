package handlers

import (
	"net/http"

	"github.com/Flook2563/test_restgolang/db"
	"github.com/Flook2563/test_restgolang/models"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db.GetDB().Create(&customer)
	c.JSON(http.StatusOK, customer)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := db.GetDB().First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := db.GetDB().First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":  "Invalid input"})
		return
	}
	db.GetDB().Save(&customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := db.GetDB().First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	db.GetDB().Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
