package main

import (
	"time"

	"github.com/tapsaca/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	config := config {
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&config)
}