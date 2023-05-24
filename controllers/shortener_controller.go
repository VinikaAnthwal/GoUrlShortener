package controllers

import (
	"go-url-shortener/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateShortURL creates a short URL from a long URL and stores it in the database
// c *gin.Context is a pointer to the Gin context
func CreateShortURL(c *gin.Context) {
    // Create a URL struct
    var url models.URL
    // Bind the JSON body to the URL struct
    c.BindJSON(&url)
    // Generate a short URL
    url.GenerateShortURL()
    // Create the URL in the database
    models.CreateURL(&url)
    // Return the short URL in the response
    c.JSON(200, gin.H{"short_url": url.ShortURL})
}

// RedirectShortURL redirects the user to the long URL associated with the short URL
// c *gin.Context is a pointer to the Gin context
func RedirectShortURL(c *gin.Context) {
    shortURL := c.Param("short_url")
    url, err := models.GetURLByShortURL(shortURL)
    if err != nil {
        c.JSON(404, gin.H{"error": "URL not found"})
        return
    }
    
    // Update statistics
    url.AccessCount++
    now := time.Now()
    url.LastAccessed = &now
    url.AccessPlace = c.ClientIP()
    models.UpdateURL(&url)
    
    c.Redirect(301, url.LongURL)
}



// GetURLStatistics returns the statistics for a specific short URL
// c *gin.Context is a pointer to the Gin context
func GetURLStatistics(c *gin.Context) {
    shortURL := c.Param("short_url")
    url, err := models.GetURLByShortURL(shortURL)
    if err != nil {
        c.JSON(404, gin.H{"error": "URL not found"})
        return
    }

    c.JSON(200, gin.H{
        "short_url":     url.ShortURL,
        "long_url":      url.LongURL,
        "access_count":  url.AccessCount,
        "last_accessed": url.LastAccessed,
        "access_place":  url.AccessPlace,
    })
}
