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

	r.GET("", api.GetInventories)
	r.GET("/:id", api.GetInventory)
	r.POST("", api.CreateInventory)
	r.PUT("/:id", api.UpdateInventory)
	r.DELETE("/:id", api.DeleteInventory)

	return r
}
