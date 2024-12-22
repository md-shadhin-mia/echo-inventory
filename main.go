package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/md-shadhin-mia/echo-inventory/controllers"
)

var db *gorm.DB
var err error

// Define the model
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Connect to SQLite database
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Automatically migrate the schema (create tables)
	db.AutoMigrate(&User{})

	// Set up Gin router
	r := gin.Default()
	

	controllers.NewCategoryController(r.Group("categories"))

	// Run the server
	r.Run(":8080")
}