package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	paumentRoutes "github.com/iamsabbiralam/stripe-with-go/pkg/payment/routes"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/ping", pingHandler)
	baseURL := router.Group("/")
	paumentRoutes.PaymentRoutes(baseURL, db)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is running"})
}
