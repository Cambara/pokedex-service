package models

import "gorm.io/gorm"

const (
	BattleTypeOffensiveEnum = "offensive"
	BattleTypeDefensiveEnum = "defensive"
)

type TypeBattleProperty struct {
	gorm.Model
	TypeId     int    `gorm:"not null"`
	Type       Type   `gorm:"foreignKey:type_id;references:id"`
	BattleType string `gorm:"type:enum('offensive', 'defensive'); not null"`
	Power      float32
}
