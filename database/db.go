package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/toluwase1/golang-inventory-system/models"
	"log"
	"os"
)

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://postgres:toluwase@localhost:5432/inventory?sslmode=disable"
	}

	//url := "postgres://postgres:toluwase@localhost:5432/inventory?sslmode=verify-full"

	db, err := gorm.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate([]models.Inventory{})
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
