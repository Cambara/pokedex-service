package database

import (
	"github.com/Cambara/pokedex-service/dto/models"
)

func Migrate() {
	db := GetInstance().Db
	db.AutoMigrate(
		models.Type{},
		models.Ability{},
		models.Move{},
		models.TypeBattleProperty{},
		models.Pokemon{},
		models.PokemonStatus{},
		models.PokemonEvolution{},
	)
}
