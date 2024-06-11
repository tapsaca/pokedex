package main

import "github.com/tapsaca/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
}

func main() {
	config := config {
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&config)
}