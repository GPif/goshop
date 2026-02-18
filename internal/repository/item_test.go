package repository

import (
	"goshop/internal/model"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	if err := db.AutoMigrate(&model.Item{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestAddToShoppingList(t *testing.T) {
	db := setupTestDB(t)

	// Test adding a single item
	err := AddToShoppingList(db, "Milk")
	if err != nil {
		t.Fatalf("AddToShoppingList failed: %v", err)
	}

	// Verify item was created
	var item model.Item
	result := db.Where("title = ?", "Milk").First(&item)
	if result.Error != nil {
		t.Fatalf("failed to find item: %v", result.Error)
	}

	if item.Title != "Milk" {
		t.Errorf("expected title 'Milk', got '%s'", item.Title)
	}
}

func TestAddToShoppingList_Duplicate(t *testing.T) {
	db := setupTestDB(t)

	// Add first item
	err := AddToShoppingList(db, "Milk")
	if err != nil {
		t.Fatalf("first AddToShoppingList failed: %v", err)
	}

	// Try to add duplicate
	err = AddToShoppingList(db, "Milk")
	if err != nil {
		t.Fatalf("second AddToShoppingList failed: %v", err)
	}

	// Should still only have one item
	var count int64
	db.Model(&model.Item{}).Count(&count)
	if count != 1 {
		t.Errorf("expected 1 item, got %d", count)
	}
}

func TestAddToShoppingList_Multiple(t *testing.T) {
	db := setupTestDB(t)

	err := AddToShoppingList(db, "Milk", "Bread", "Eggs")
	if err != nil {
		t.Fatalf("AddToShoppingList failed: %v", err)
	}

	var count int64
	db.Model(&model.Item{}).Count(&count)
	if count != 3 {
		t.Errorf("expected 3 items, got %d", count)
	}
}

func TestRemoveFromShoppingList(t *testing.T) {
	db := setupTestDB(t)

	err := AddToShoppingList(db, "Milk", "Bread", "Eggs")
	if err != nil {
		t.Fatalf("AddToShoppingList failed: %v", err)
	}

	err = RemoveFromShoppingList(db, 1)
	if err != nil {
		t.Fatalf("RemoveFromShoppingList failed: %v", err)
	}

	var count int64
	db.Model(&model.Item{}).Count(&count)
	if count != 2 {
		t.Errorf("expected 2 items, got %d", count)
	}
}

func TestRemoveFromShoppingList_NonExistent(t *testing.T) {
	db := setupTestDB(t)

	err := AddToShoppingList(db, "Milk", "Bread", "Eggs")
	if err != nil {
		t.Fatalf("AddToShoppingList failed: %v", err)
	}

	err = RemoveFromShoppingList(db, 4)
	if err != nil {
		t.Fatalf("RemoveFromShoppingList failed: %v", err)
	}

	var count int64
	db.Model(&model.Item{}).Count(&count)
	if count != 3 {
		t.Errorf("expected 3 items, got %d", count)
	}
}

func TestRemoveFromShoppingList_Multiple(t *testing.T) {
	db := setupTestDB(t)

	err := AddToShoppingList(db, "Milk", "Bread", "Eggs")
	if err != nil {
		t.Fatalf("AddToShoppingList failed: %v", err)
	}

	err = RemoveFromShoppingList(db, 1, 2)
	if err != nil {
		t.Fatalf("RemoveFromShoppingList failed: %v", err)
	}

	var count int64
	db.Model(&model.Item{}).Count(&count)
	if count != 1 {
		t.Errorf("expected 1 item, got %d", count)
	}
}
