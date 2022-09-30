package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	ID     int
	Number int
	Name   string
}
