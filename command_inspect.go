package main

import (
	"errors"
	"fmt"
)

func callbackInspect(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := config.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught this pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}