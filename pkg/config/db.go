package config

import (
	"fmt"
	"log"
	"os"

	paymentModels "github.com/iamsabbiralam/stripe-with-go/pkg/payment/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func Migrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(
		&paymentModels.Payment{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed")
}

func CloseDB(DB *gorm.DB) {
	db, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get DB from GORM: ", err)
	}

	db.Close()
}
