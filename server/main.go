package main

import (
	"log"
	"os"
	"template-starter-mysql/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Load konfigurasi dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup rute-rute API
	routes.SetupRoutes(app)

	// Mulai server
	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
