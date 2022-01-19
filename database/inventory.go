package database

import (
	"github.com/jinzhu/gorm"
	"github.com/toluwase1/golang-inventory-system/models"
)

func GetInventories(db *gorm.DB) ([]models.Inventory, error) {
	inventories := []models.Inventory{}
	query := db.Select("inventories.*").
		Group("inventories.id")
	if err := query.Find(&inventories).Error; err != nil {
		return inventories, err
	}

	return inventories, nil
}

func GetInventoryByID(id string, db *gorm.DB) (models.Inventory, bool, error) {
	inventory := models.Inventory{}

	query := db.Select("inventories.*")
	query = query.Group("inventories.id")
	err := query.Where("inventories.id = ?", id).First(&inventory).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return inventory, false, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return inventory, false, nil
	}
	return inventory, true, nil
}

func DeleteInventory(id string, db *gorm.DB) error {
	var inv models.Inventory
	if err := db.Where("id = ? ", id).Delete(&inv).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInventory(db *gorm.DB, inv *models.Inventory) error {
	if err := db.Save(&inv).Error; err != nil {
		return err
	}
	return nil
}
