package main

import (
	"fmt"

	"github.com/bruttins/pokedexGo/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	areaName := args[0]
	locationInfo, err := pokeapi.GetLocationInfo(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Printf("Found Pokemon:\n")
	for _, pokemonName := range locationInfo.Encounters {
		fmt.Println(pokemonName.Pokemon.Name)
	}
	return nil
}
