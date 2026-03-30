package main

import (
	"fmt"

	"github.com/bruttins/pokedexGo/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
	fullUrl := cfg.Next
	locations, err := pokeapi.GetLocation(fullUrl)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.Previous = locations.Previous
	cfg.Next = locations.Next
	return nil
}
