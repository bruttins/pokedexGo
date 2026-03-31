package main

import (
	"fmt"
	"math/rand"

	"github.com/bruttins/pokedexGo/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	pokemonName := args[0]
	pokemonInfo, err := pokeapi.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	pokeExp := pokemonInfo.BaseExperience
	k := 80.00
	catchChance := k / (k + float64(pokeExp))

	if rand.Float64() < (catchChance) {
		fmt.Println(pokemonInfo.Name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.Pokedex[pokemonName] = pokemonInfo
	} else {
		fmt.Println(pokemonInfo.Name + " escaped!")
	}
	return nil
}
