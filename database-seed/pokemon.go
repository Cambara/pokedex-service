package databaseseed

import (
	"fmt"

	"github.com/Cambara/pokedex-service/database-seed/data"
)

func importPokemons() {
	pokemons := data.GetPokemonData01()
	fmt.Println(len(pokemons))
	/*
		for _, pokemon := range pokemons {
			fmt.Println(pokemon)
		}
	*/
}
