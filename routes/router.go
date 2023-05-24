package routes

import (
	"go-url-shortener/controllers"
	"go-url-shortener/models"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes the routes for the application
// It takes a pointer to a gin.Engine as an argument and sets up the routes
// It also initializes the database connection
func InitializeRoutes(router *gin.Engine) {
    models.InitDB()

    router.POST("/shorten", controllers.CreateShortURL)
    router.GET("/:short_url", controllers.RedirectShortURL)
    router.GET("/api/:short_url/stats", controllers.GetURLStatistics)

    router.Use(func(c *gin.Context) {
        c.JSON(404, gin.H{"error": "Not Found"})
    })

    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{"error": "Not Found"})
    })
}
