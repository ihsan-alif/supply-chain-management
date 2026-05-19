package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	DB = database

	fmt.Println("Database connected")

	runMigration()
}

func runMigration() {

	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration success")
}
