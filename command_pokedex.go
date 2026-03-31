package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for i := range cfg.Pokedex {
		fmt.Println("- " + cfg.Pokedex[i].Name)
	}
	return nil
}
