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