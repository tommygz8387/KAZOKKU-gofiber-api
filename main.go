package main

import (
	"fmt"
	"log"
	"os"
	"v1/app/database"
	"v1/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database connection
    _, err = database.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer database.CloseDB()

    // Create Fiber app
    app := fiber.New()

    // Initialize routes
    routes.SetupRoutes(app)

    // Start server
    port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
    log.Fatal(app.Listen(port))
}
