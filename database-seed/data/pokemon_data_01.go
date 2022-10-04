package data

import (
	"os"

	"github.com/gocarina/gocsv"
)

type PokemonData01Dto struct {
	Number    int    `csv:"pokedex_number"`
	Name      string `csv:"name"`
	Abilities string `csv:"abilities"`
	Type1     string `csv:"type1"`
	Type2     string `csv:"type2"`
}

func GetPokemonData01() []*PokemonData01Dto {
	file, err := os.OpenFile("database-seed/data/pokemon-data-01.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	pokemons := []*PokemonData01Dto{}

	err = gocsv.UnmarshalFile(file, &pokemons)

	if err != nil {
		panic(err)
	}

	return pokemons
}
