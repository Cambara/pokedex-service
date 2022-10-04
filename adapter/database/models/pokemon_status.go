package models

import "gorm.io/gorm"

type PokemonStatus struct {
	gorm.Model
	Hp             int `gorm:"not null"`
	Attack         int `gorm:"not null"`
	Defense        int `gorm:"not null"`
	SpecialAttack  int `gorm:"not null"`
	SpecialDefense int `gorm:"not null"`
	Speed          int `gorm:"not null"`
}
