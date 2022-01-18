package models

type Inventory struct {
	gorm.Model
	Name        string
	Description string
	Price       int
	Quantity    int
}
