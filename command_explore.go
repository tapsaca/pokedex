package main

import (
	"errors"
	"fmt"
)

func callbackExplore(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]
	locationArea, error := config.pokeapiClient.GetLocationArea(locationAreaName)
	if error != nil {
		return error
	}
	fmt.Println("")
	fmt.Printf("Pokemon found in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}