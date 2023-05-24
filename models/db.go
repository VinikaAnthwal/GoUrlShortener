package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	// Get the database connection string from the environment variable.
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	// Connect to the database using the gorm library.
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Automigrate the URL struct.
	db.AutoMigrate(&URL{})
}
