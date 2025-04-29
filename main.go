package main

import (
	"log"
	"net/http"

	"github.com/iamsabbiralam/stripe-with-go/pkg/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	configure := cors.DefaultConfig()
	configure.AllowOrigins = []string{
		"*",
	}
	configure.AllowCredentials = true

	// allow all headers
	configure.AllowHeaders = []string{"*"}
	router.Use(cors.New(configure))

	DB, err = config.InitDB()
	if err != nil {
		log.Fatal("Could not connect to the database", err)
		return
	}
	defer config.CloseDB(DB)

	// config.Migrate(DB)
	// Root endpoint handling for GET and HEAD
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is running..."})
	})
	router.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	config.InitializeRoutes(router, DB)
	// utils.RegisterEndpoints(router, DB)
	config.InitStripe()

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Could not start the server", err)
		return
	}
}
