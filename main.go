package main

import (
	"inventory-management-golang/database"
	"inventory-management-golang/handlers"

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
		transactions := api.Group("/transactions")
		{
			transactions.GET("", handlers.GetTransactions)
			transactions.POST("", handlers.CreateTransaction)
		}

		// Supplier routes 
		suppliers := api.Group("/suppliers")
		{
			suppliers.GET("", handlers.GetSuppliers)
			suppliers.POST("", handlers.CreateSupplier)
		}

		// Alert routes 
		alerts := api.Group("/alerts")
		{
			alerts.GET("", handlers.GetStockAlerts)
			alerts.PUT("/:id/resolve", handlers.ResolveAlert)
		}

		// Dashboard route 
		api.GET("/dashboard", handlers.GetDashboardStats)
	}

	// Start server 
	r.Run(":8080")
}