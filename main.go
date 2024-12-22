package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/md-shadhin-mia/echo-inventory/controllers"
	"github.com/md-shadhin-mia/echo-inventory/models"
)

var db *gorm.DB
var err error

// Define the model
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DB *gorm.DB

func main() {
	// Connect to SQLite database
	// db := initialize.DB

	// Automatically migrate the schema (create tables)
	// db.AutoMigrate(&User{})
	//if args have --migrate then automgrate

	if len(os.Args) > 1 && os.Args[1] == "--migrate" {
		db, err = gorm.Open("sqlite3", "./db.sqlite")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
		db.AutoMigrate(&models.Category{}, &models.ProductType{})
		return
	}

	// Set up Gin router
	r := gin.Default()

	controllers.NewCategoryController(r.Group("categories"))
	controllers.NewProductTypeController(r.Group("product-types"))

	// Run the server
	r.Run(":8080")
}
