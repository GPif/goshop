package model

import "time"

type Item struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Completed bool      `gorm:"default:false" json:"completed"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
