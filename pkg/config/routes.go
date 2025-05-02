package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	paumentRoutes "github.com/iamsabbiralam/stripe-with-go/pkg/payment/routes"
)

// InitializeRoutes initializes all routes for the application
func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	// Ping Server
	router.GET("/ping", pingHandler)

	// Create a router group with the base URL "/"
	baseURL := router.Group("/")
	paumentRoutes.PaymentRoutes(baseURL, db)
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "API is running"})
}
