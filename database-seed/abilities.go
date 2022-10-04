package databaseseed

import (
	"strings"

	"github.com/Cambara/pokedex-service/adapter/database"
	"github.com/Cambara/pokedex-service/adapter/database/models"
	"github.com/Cambara/pokedex-service/database-seed/data"
	"golang.org/x/exp/maps"
)

func importAbilities() {
	pokemons := data.GetPokemonData01()
	abilitiesMap := make(map[string]bool)
	for _, pokemon := range pokemons {
		pokemonAbilities := ConvertStringInArray(pokemon.Abilities)
		for _, pokemonAbility := range pokemonAbilities {
			key := strings.ToLower(pokemonAbility)
			_, value := abilitiesMap[key]
			if !value {
				abilitiesMap[key] = true
			}
		}
	}
	abilities := maps.Keys(abilitiesMap)

	db := database.GetInstance()

	for _, ability := range abilities {
		db.Db.Create(&models.Ability{
			Name: ability,
		})
	}
}
