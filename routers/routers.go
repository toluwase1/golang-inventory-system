package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/golang-inventory-system/database"
	"github.com/toluwase1/golang-inventory-system/handlers"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	router := r.Group("/inventories")

	router.GET("/all", api.GetInventories)
	router.GET("/:id", api.GetInventory)
	router.POST("/create", api.CreateInventory)
	router.PUT("/:id", api.UpdateInventory)
	router.DELETE("/:id", api.DeleteInventory)

	return r
}
