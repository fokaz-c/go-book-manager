package config

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	var err error
	dbPath := "db/gorm.db"

	// Ensure the directory exists
	dir := filepath.Dir(dbPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("Directory does not exist: %s, creating it...", dir)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Panicf("failed to create directory: %v", err)
		}
	}

	log.Printf("Attempting to open database file at: %s", dbPath)
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully")
}

func GetDB() *gorm.DB {
	return db
}
