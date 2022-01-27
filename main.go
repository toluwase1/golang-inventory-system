package main

import (
	"fmt"
	"github.com/toluwase1/golang-inventory-system/database"
	_ "github.com/toluwase1/golang-inventory-system/docs"
	"github.com/toluwase1/golang-inventory-system/routers"
	"log"
	"os"
)

// @title           Shopify Inventory Management System API
// @version         1.0
// @description     This is a  Inventory Management System API for Shopify.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
func main() {
	database.Setup()
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT = ":8080"
	}
	r := routers.Setup()
	if err := r.Run(PORT); err != nil {
		log.Fatal(err)
	}

}
