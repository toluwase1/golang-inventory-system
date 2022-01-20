package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/toluwase1/golang-inventory-system/database"
	"github.com/toluwase1/golang-inventory-system/models"
	"net/http"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetInventories(c *gin.Context) {
	inventories, err := database.GetInventories(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, inventories)
}
func (a *APIEnv) Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the API",
	})
}

func (a *APIEnv) GetInventory(c *gin.Context) {
	id := c.Params.ByName("id")
	inv, exists, err := database.GetInventoryByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "there is no inventory in db")
		return
	}

	c.JSON(http.StatusOK, inv)
}

func (a *APIEnv) CreateInventory(c *gin.Context) {
	inv := models.Inventory{}
	err := c.BindJSON(&inv)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&inv).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, inv)
}

func (a *APIEnv) DeleteInventory(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetInventoryByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	err = database.DeleteInventory(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "inventory record deleted successfully")
}

func (a *APIEnv) UpdateInventory(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetInventoryByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "record not exists")
		return
	}

	updatedInventory := models.Inventory{}
	err = c.BindJSON(&updatedInventory)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.UpdateInventory(a.DB, &updatedInventory); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.GetInventory(c)
}
