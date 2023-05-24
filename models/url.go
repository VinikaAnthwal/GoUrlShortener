package models

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db is a global variable used to store the database connection
var db *gorm.DB

// URL is a struct that stores the information of a URL
// URL is a struct that stores the information of a URL
type URL struct {
    gorm.Model
    LongURL      string     `json:"long_url" gorm:"unique"`
    ShortURL     string     `json:"short_url" gorm:"unique"`
    AccessCount  uint       `json:"access_count"`
    LastAccessed *time.Time `json:"last_accessed"`
    AccessPlace  string     `json:"access_place"`
}


// GenerateShortURL is a method used to generate a random short URL
// It takes no parameters and returns nothing
func (u *URL) GenerateShortURL() {
 rand.Seed(time.Now().UnixNano())
 var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
 b := make([]rune, 6)
 for i := range b {
  b[i] = letterRunes[rand.Intn(len(letterRunes))]
 }
 u.ShortURL = string(b)
}

// CreateURL is a method used to create a new URL in the database
// It takes a pointer to a URL struct as a parameter and returns nothing
func CreateURL(url *URL) {
 db.Create(&url)
}

// GetURLByShortURL is a method used to get a URL from the database by its short URL
// It takes a string as a parameter and returns a URL struct and an error
func GetURLByShortURL(shortURL string) (URL, error) {
 var url URL
 if err := db.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
  return url, err
 }
 return url, nil
}

//updates a URL in the database
func UpdateURL(url *URL) error {
    result := db.Save(url)
    return result.Error
}

func CloseDB() {
	// Close the database connection.
	db.Close()
}