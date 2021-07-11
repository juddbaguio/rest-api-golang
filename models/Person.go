package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model

	Name string `json:"name"`
	Email string `json:"email"`
	Books []Book `json:"books"`
}