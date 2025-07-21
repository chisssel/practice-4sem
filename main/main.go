package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"practice-4sem/config"
	"practice-4sem/models"
	"practice-4sem/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := config.BuildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		log.Fatal("Failed to migrate database")
	}

	r := routes.SetupRoutes(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5438"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(r.Run(":" + port))
}
