package models

import "gorm.io/gorm"

type PokemonEvolution struct {
	gorm.Model
	EvoluteId              int     `gorm:"not null"`
	Evolute                Pokemon `gorm:"foreignKey:evolute_id;references:id"`
	Stage                  int
	Level                  int
	Item                   string
	RequirementDescription string
}
