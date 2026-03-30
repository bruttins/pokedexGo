package main

import (
	"fmt"

	"github.com/bruttins/pokedexGo/internal/pokeapi"
)

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	fullUrl := cfg.Previous
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
