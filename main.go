package main

import (
    "log"
    "os"

    "crud-api/internal/database"
    "crud-api/internal/routes"

    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    if err := database.Connect(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer database.Close()

    // Setup routes
    r := routes.SetupRoutes()

    // Get port from environment
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
