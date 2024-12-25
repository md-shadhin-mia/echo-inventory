package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/md-shadhin-mia/echo-inventory/controllers"
	"github.com/md-shadhin-mia/echo-inventory/initialize"
	"github.com/md-shadhin-mia/echo-inventory/models"
)

var db *gorm.DB

// Define the model
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Connect to SQLite database
	// db := initialize.DB
	log.Println(" Starting server on port 8080")
	// Automatically migrate the schema (create tables)
	// db.AutoMigrate(&User{})
	//if args have --migrate then automgrate
	for _, arg := range os.Args {
		if arg == "--migrate" {
			db = initialize.DB
			db.AutoMigrate(&models.Category{}, &models.ProductType{})
			log.Println(" Migrated successfully")
		}
	}
	defer initialize.DB.Close()
	// Set up Gin router
	r := gin.Default()
	r.GET("/heath", health)
	//controllers
	controllers.NewCategoryController(r.Group("categories"))
	controllers.NewProductTypeController(r.Group("product-types"))

	// Run the server
	r.Run(":8080")
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Application is running",
		"status":  "OK",
	})
}
