package main

import (
	"example/web-service-gin/api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=root sslmode=disable")
	if err != nil {
		panic("Failed to ping database")
	}
	// defer db.Close()
	// panic("berhasil")

	router := gin.Default()

	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Auto migrate the model schema
	db.AutoMigrate(&models.Users{})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Bg Rio!",
		})
	})
}
