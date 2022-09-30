package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Number    int           `gorm:"not null"`
	Name      string        `gorm:"not null"`
	Types     []Type        `gorm:"many2many:pokemon_has_type"`
	Abilities []Ability     `gorm:"many2many:pokemon_has_ability"`
	Status    PokemonStatus `gorm:"foreignKey:id"`
}
