package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bruttins/pokedexGo/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Pokedex: map[string]pokeapi.PokemonInfo{},
	}
	for {
		fmt.Printf("Pokedex > ")
		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		cleanedWords := cleanInput(text)
		if len(cleanedWords) == 0 {
			continue
		}
		cmd := cleanedWords[0]
		value, exists := getCommands()[cmd]
		if exists {
			err := value.callback(cfg, cleanedWords[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(input string) []string {
	result := []string{}
	splitArray := strings.Fields(input)
	for _, word := range splitArray {
		cleaned := strings.ToLower(word)
		result = append(result, cleaned)
	}
	return result
}

type config struct {
	Next     *string
	Previous *string
	Pokedex  map[string]pokeapi.PokemonInfo
}
