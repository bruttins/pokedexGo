package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bruttins/pokedexGo/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Second)

func GetLocation(pageURL *string) (Response, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	locations := Response{}
	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return Response{}, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Response{}, fmt.Errorf("error: %w", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return Response{}, err
		}
		cache.Add(url, data)
		if err := json.Unmarshal(data, &locations); err != nil {
			return Response{}, err
		}
	}
	return locations, nil
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Response struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

func GetLocationInfo(name string) (LocationInfo, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	locationInfo := LocationInfo{}
	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &locationInfo); err != nil {
			return LocationInfo{}, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationInfo{}, fmt.Errorf("error: %w", err)
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationInfo{}, err
		}
		cache.Add(url, data)
		if err := json.Unmarshal(data, &locationInfo); err != nil {
			return LocationInfo{}, err
		}
	}
	return locationInfo, nil
}

type LocationInfo struct {
	Id         int                 `json:"id"`
	Name       string              `json:"name"`
	Location   Location            `json:"location"`
	Encounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetPokemonInfo(name string) (PokemonInfo, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	pokemonInfo := PokemonInfo{}
	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemonInfo); err != nil {
			return PokemonInfo{}, err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return PokemonInfo{}, fmt.Errorf("error: %w", err)
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return PokemonInfo{}, err
		}
		cache.Add(url, data)
		if err := json.Unmarshal(data, &pokemonInfo); err != nil {
			return PokemonInfo{}, err
		}
	}
	return pokemonInfo, nil
}

type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
}
