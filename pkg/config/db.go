package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

	// Create the connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT,
	)
	// dsn := "host=propteq-dev.cluster-c74yw6gco1oh.ap-southeast-2.rds.amazonaws.com user=postgres password=]k(C[tu70Y0_UXrA{$iP2Eyd6|h3 dbname=propteq port=5432"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return nil, err
	}
	fmt.Println("Connected to the Database")
	return DB, nil
}

func Migrate(db *gorm.DB) {
	// Enable the "uuid-ossp" extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(
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
