package entities

import "gorm.io/gorm"

type Student struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	gorm.Model
}
