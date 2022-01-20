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
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT = ":8080"
	}
	r := routers.Setup()
	if err := r.Run(PORT); err != nil {
		log.Fatal(err)
	}

}
