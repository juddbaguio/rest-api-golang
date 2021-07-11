package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title string
	Author string
	CallNumber int `gorm:"unique_index"`
	PersonID int
}