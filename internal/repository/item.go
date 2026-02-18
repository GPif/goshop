package repository

import (
	"errors"
	"goshop/internal/model"

	"gorm.io/gorm"
)

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
