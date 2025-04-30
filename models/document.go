package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Content    string `json:"content"`
}