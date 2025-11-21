package handlers

import (
	"inventory-management-golang/database"
	"inventory-management-golang/models"
	"net/http"
	"strconv"

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
	id := c.Param("id")
	var product models.Product 
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return 
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c*gin.Context) {
	var product models.Product 
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check stock alert after creation 
	CheckStockAlert(&product)

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product 
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	// Check stock alert after creation 
	CheckStockAlert(&product)

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// Transaction Handlers 
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return 
	}

	// Get product 
	var product models.Product 
	if err := database.DB.First(&product, transaction.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return 
	}

	// Update stock based on transaction type
	if transaction.Type == "in" {
		product.Stock += transaction.Quantity
	} else if transaction.Type == "out" {
		if product.Stock < transaction.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufiicient Stock"})
			return 
		}
		product.Stock -= transaction.Quantity
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction type"})
		return 
	}
	
	// Start transaction 
	tx := database.DB.Begin()

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	// Check stock alert 
	CheckStockAlert(&product)

	tx.Commit()
	c.JSON(http.StatusCreated, transaction)
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction 
	if err := database.DB.Preload("Product").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, transactions)
}

// Supplier Handlers 
func GetSuppliers(c *gin.Context) {
	var suppliers []models.Supplier
	if err := database.DB.Find(&suppliers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, suppliers)
}

func CreateSupplier(c *gin.Context) {
	var supplier models.Supplier 
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := database.DB.Create(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, supplier)
}

// Stock Alert Handlers 
func GetStockAlerts(c *gin.Context) {
	var alerts []models.StockAlert
	if err := database.DB.Preload("Product").Where("is_resolved = ?", false).Find(&alerts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusOK, alerts)
}

func ResolveAlert(c *gin.Context) {
	id := c.Param("id")
	var alert models.StockAlert
	if err := database.DB.First(&alert, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return 
	}

	alert.IsResolved = true 
	if err := database.DB.Save(&alert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, alert)
}

// Helper function to check stock alerts 
func CheckStockAlert(product *models.Product) {
	if product.Stock <= product.MinStock {
		// Check if alert already exists
		var existingAlert models.StockAlert 
		if err := database.DB.Where("product_id = ? AND is_resolved = ?", product.ID, false).First(&existingAlert).Error; err != nil {
			// Create new alert
			alert := models.StockAlert {
				ProductID: product.ID,
				Message: "Stock rendah untuk produk " + product.Name + " . Stok tersisa: " + strconv.Itoa(product.Stock),
			}
			database.DB.Create(&alert)
		} 
	}
}

// Dashbaord Statistics 
