package handlers

import (
	"inventory-management-golang/database"
	"inventory-management-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Product Handlers
func GetProducts(c *gin.Context) {
	var products []models.Product 
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Params("id")
	var product models.Product 
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return 
	}
	c.JSON(http.StatusOK, product)
}

