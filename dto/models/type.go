package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Name string `gorm:"not null"`
}