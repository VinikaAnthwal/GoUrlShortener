package main

import (
	"go-url-shortener/models"
	"go-url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    // Create a new Gin router.
    r := gin.Default()

    // Initialize the database connection.
    models.InitDB()

    // Initialize the routes.
    routes.InitializeRoutes(r)

    // Defer the closing of the database connection.
    defer models.CloseDB()

    // Run the router.
    r.Run()
}
