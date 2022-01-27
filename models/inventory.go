package models

import "github.com/jinzhu/gorm"

type Inventory struct {
	gorm.Model
	Name        string
	Description string
	Price       int
	Quantity    int
}

//type InventoryResponse struct {
//}
//
//type ErrorResponse struct {
//}
