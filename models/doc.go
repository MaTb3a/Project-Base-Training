package models

import (
	"time"

	"gorm.io/gorm"
)

type DocumentFromOrm struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Title     string
	Author    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Document struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Author    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
