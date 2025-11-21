package main

import (
	"inventory-management-golang/database"
	"inventory-management-golang/handlers"
	"inventory-management/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database 
	database.InitDatabase();

	// Create gin router 
	r := gin.Default()

	// CORS middleware 
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return 
		}

		c.Next()
	})

	// Routes 
	api := r.Group("/api")
	{
		// Product routes 
		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.POST("", handlers.CreateProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
		}

		// Transaction routes 
		
	}

}