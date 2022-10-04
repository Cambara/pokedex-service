package models

import "gorm.io/gorm"

type Move struct {
	gorm.Model
	Name     string `gorm:"not null"`
	TypeId   int    `gorm:"not null"`
	Type     Type   `gorm:"foreignKey:type_id;references:id"`
	Category string `gorm:"not null"`
	Contest  string `gorm:"not null"`
	PP       int    `gorm:"not null"`
	Power    int    `gorm:"not null"`
	Accuracy int
}
