package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Number     int                `gorm:"not null"`
	Name       string             `gorm:"not null"`
	Status     PokemonStatus      `gorm:"foreignKey:id"`
	Types      []Type             `gorm:"many2many:pokemon_has_type"`
	Abilities  []Ability          `gorm:"many2many:pokemon_has_ability"`
	Evolutions []PokemonEvolution `gorm:"many2many:pokemon_has_evolution"`
}
