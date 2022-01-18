package database

import (
	"github.com/jinzhu/gorm"
	"github.com/toluwase1/golang-inventory-system/models"
	"log"
)

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
	host := "host"
	port := "port"
	dbname := "dbname"
	user := "user"
	password := "password"

	db, err := gorm.Open("postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbname+
			" sslmode=disable password="+password)

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
