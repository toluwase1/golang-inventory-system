package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/toluwase1/golang-inventory-system/models"
	"log"
)

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
	//host := "host"
	//port := "8080"
	//dbname := "inventory"
	//user := "postgres"
	//password := "toluwase"
	url := "postgres://postgres:toluwase@localhost:5432/inventory?sslmode=disable"
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
