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

// ShowInventory godoc
// @Summary      Show an inventory
// @Description  get string by ID
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "Inventory ID"
// @Success      200  {object}  models.Inventory
// @Failure      400  {object}  docs.HTTPError
// @Failure      404  {object}  docs.HTTPError
// @Failure      500  {object}  docs.HTTPError
// @Router       /inventories/ [get]
func (a *APIEnv) GetInventories(c *gin.Context) {
	inventories, err := database.GetInventories(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, inventories)
}

// ShowInventory godoc
// @Summary      Welcome page
// @Description  welcome page
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "Inventory ID"
// @Success      200  {object}  models.Inventory
// @Failure      400  {object}  docs.HTTPError
// @Failure      404  {object}  docs.HTTPError
// @Failure      500  {object}  docs.HTTPError
// @Router       / [get]
func (a *APIEnv) Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Welcome Message":     "Welcome to the Shopify Inventory System",
		"Instruction":         "Send request on postman",
		"Get All Inventories": "http://inventory-manager-shopify-task.herokuapp.com/inventories/all",
		"Get One Inventory":   "http://inventory-manager-shopify-task.herokuapp.com/inventories/:id",
		"Create Inventory":    "http://inventory-manager-shopify-task.herokuapp.com/inventories/create",
		"Delete Inventory":    "http://inventory-manager-shopify-task.herokuapp.com/inventories/:id",
		"Update Inventory":    "http://inventory-manager-shopify-task.herokuapp.com/inventories/:id",

		"MODELS": " Name  string, Description string, Price int, Quantity int ",
	})
}

// ShowInventory godoc
// @Summary      Show an inventory
// @Description  get string by ID
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "Inventory ID"
// @Success      200  {object}  models.Inventory
// @Failure      400  {object}  docs.HTTPError
// @Failure      404  {object}  docs.HTTPError
// @Failure      500  {object}  docs.HTTPError
// @Router       /inventories/{id} [get]
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

// ShowInventory godoc
// @Summary      Create an inventory
// @Description  create inventory
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "Inventory ID"
// @Success      200  {object}  models.Inventory
// @Failure      400  {object}  docs.HTTPError
// @Failure      404  {object}  docs.HTTPError
// @Failure      500  {object}  docs.HTTPError
// @Router       /inventories/create/ [post]
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

// ShowInventory godoc
// @Summary      Delete an inventory
// @Description  delete inventory by ID
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "Inventory ID"
// @Success      200  {object}  models.Inventory
// @Failure      400  {object}  docs.HTTPError
// @Failure      404  {object}  docs.HTTPError
// @Failure      500  {object}  docs.HTTPError
// @Router       /inventories/{id}/ [delete]
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

// ShowInventory godoc
// @Summary      update an inventory
// @Description  Update inventory
// @Tags         accounts
// @Produce      json
// @Param        id   path      int  true  "Inventory ID"
// @Success      200  {object}  models.Inventory
// @Failure      400  {object}  docs.HTTPError
// @Failure      404  {object}  docs.HTTPError
// @Failure      500  {object}  docs.HTTPError
// @Router       /inventories/{id} [patch]
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
