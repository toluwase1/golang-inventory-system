package main

import (
	"fmt"
	"github.com/toluwase1/golang-inventory-system/database"
	"github.com/toluwase1/golang-inventory-system/routers"
	"log"
	"os"
)

func main() {
	database.Setup()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == ":" {
		port = ":8080"
	}
	r := routers.Setup()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}

}
