package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/toluwase1/golang-inventory-system/handlers"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/toluwase1/golang-inventory-system/database"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	r.GET("/", api.Welcome)
	router := r.Group("/inventories")

	router.GET("/all", api.GetInventories)
	router.GET("/:id", api.GetInventory)
	router.POST("/create", api.CreateInventory)
	router.PUT("/:id", api.UpdateInventory)
	router.DELETE("/:id", api.DeleteInventory)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
