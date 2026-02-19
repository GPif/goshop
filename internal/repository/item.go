package repository

import (
	"errors"
	"fmt"
	"goshop/internal/model"

	"gorm.io/gorm"
)

func ViewShoppingList(db *gorm.DB) error {
	var items []model.Item
	res := db.Find(&items)
	if res.Error != nil {
		return res.Error
	}
	fmt.Println("Shopping List:")
	for _, item := range items {
		fmt.Printf("%d - %s\n", item.ID, item.Title)
	}
	return nil
}

func AddToShoppingList(db *gorm.DB, names ...string) error {
	for _, name := range names {
		item := model.Item{Title: name}

		var existingItem model.Item
		exists := db.Where("title = ?", name).First(&existingItem)

		if errors.Is(exists.Error, gorm.ErrRecordNotFound) {
			res := db.Create(&item)
			if res.Error != nil {
				return res.Error
			}
		} else if exists.Error != nil {
			return exists.Error
		}
	}
	return nil
}

func RemoveFromShoppingList(db *gorm.DB, ids ...int) error {
	var item model.Item
	res := db.Where("id in (?)", ids).Delete(&item)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func EditShoppingList(db *gorm.DB, itemID int, newName string) error {
	var item model.Item
	res := db.Where("id = ?", itemID).First(&item)
	if res.Error != nil {
		return res.Error
	}

	item.Title = newName
	res = db.Save(&item)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
