package initialize

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	// Initialize the logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(ioutil.Discard)
	//database

	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	DB = db
	// defer db.Close()
}
