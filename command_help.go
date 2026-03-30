package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range getCommands() {
		cmdName := value.name
		cmdDescription := value.description
		fmt.Printf("%s: %s\n", cmdName, cmdDescription)
	}
	fmt.Println()
	return nil
}
