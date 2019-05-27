package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Used for talking to sqlite databases
)

// Setup initializes the database and returns a pointer to it for other parts of the application
func Setup() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.LogMode(true)
	db.AutoMigrate(&Plugin{}, &Machine{}, &Server{})

	return db
}
