package databaseseed

import (
	"strings"

	"github.com/Cambara/pokedex-service/adapter/database"
	"github.com/Cambara/pokedex-service/adapter/database/models"
	"github.com/Cambara/pokedex-service/database-seed/data"
	"golang.org/x/exp/maps"
)

func importAbilities() {
	db := database.GetInstance()

	var result []*models.Ability
	var count int64
	db.Db.Model(&result).Count(&count)

	if count > 0 {
		return
	}

	pokemons := data.GetPokemonData01()
	abilitiesMap := make(map[string]bool)
	for _, pokemon := range pokemons {

		for _, pokemonAbility := range pokemon.Abilities {
			key := strings.ToLower(pokemonAbility)
			_, value := abilitiesMap[key]
			if !value {
				abilitiesMap[key] = true
			}
		}
	}
	abilities := maps.Keys(abilitiesMap)

	for _, ability := range abilities {
		db.Db.Create(&models.Ability{
			Name: ability,
		})
	}
}
