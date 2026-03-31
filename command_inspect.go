package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if _, ok := cfg.Pokedex[args[0]]; !ok {
		fmt.Println("You have not caught that pokemon.")
	} else {
		pokemon := args[0]
		fmt.Printf("Name: %s\n", cfg.Pokedex[pokemon].Name)
		fmt.Printf("Height: %v\n", cfg.Pokedex[pokemon].Height)
		fmt.Printf("Weight: %v\n", cfg.Pokedex[pokemon].Weight)
	}
	return nil
}
