package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	pokemonName := args[0]
	pokemon, error := config.pokeapiClient.GetPokemon(pokemonName)
	if error != nil {
		return error
	}
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	config.caughtPokemon[pokemonName] = pokemon
	fmt.Println("")
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("")
	return nil
}