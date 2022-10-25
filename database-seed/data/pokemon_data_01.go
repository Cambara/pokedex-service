package data

import (
	"os"
	"strings"

	"github.com/Cambara/pokedex-service/database-seed/helper"
	"github.com/gocarina/gocsv"
	"github.com/gosimple/slug"
)

type pokemonData01CSV struct {
	Number    int    `csv:"pokedex_number"`
	Name      string `csv:"name"`
	Abilities string `csv:"abilities"`
	Type1     string `csv:"type1"`
	Type2     string `csv:"type2"`
}

type PokemonData01Dto struct {
	Number    int
	Name      string
	Slug      string
	Abilities []string
	Types     []string
}

var pokemons []PokemonData01Dto

func GetPokemonData01() []PokemonData01Dto {
	if pokemons != nil {
		return pokemons
	}

	file, err := os.OpenFile("database-seed/data/pokemon-data-01.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	csvPokemons := []*pokemonData01CSV{}

	err = gocsv.UnmarshalFile(file, &csvPokemons)

	if err != nil {
		panic(err)
	}

	pokemons = []PokemonData01Dto{}

	for _, csvPokemon := range csvPokemons {
		abilities := helper.ConvertStringInArray(csvPokemon.Abilities)
		types := make([]string, 0)

		if csvPokemon.Type1 != "" {
			types = append(types, strings.TrimSpace(csvPokemon.Type1))
		}

		if csvPokemon.Type2 != "" {
			types = append(types, strings.TrimSpace(csvPokemon.Type2))
		}

		pokemon := PokemonData01Dto{
			Number:    csvPokemon.Number,
			Name:      strings.TrimSpace(csvPokemon.Name),
			Slug:      slug.Make(strings.TrimSpace(csvPokemon.Name)),
			Abilities: abilities,
			Types:     types,
		}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}
