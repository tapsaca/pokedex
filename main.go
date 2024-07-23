package main

import (
	"time"

	"github.com/tapsaca/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
}

func main() {
	config := config {
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&config)
}