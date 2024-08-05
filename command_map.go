package main

import (
	"errors"
	"fmt"
)

func callbackMap(config *config, args ...string) error {
	response, error := config.pokeapiClient.ListLocationAreas(config.nextLocationAreaURL)
	if error != nil {
		return error
	}
	fmt.Println("")
	fmt.Println("Location areas:")
	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	config.nextLocationAreaURL = response.Next
	config.previousLocationAreaURL = response.Previous
	return nil
}

func callbackMapb(config *config, args ...string) error {
	if config.previousLocationAreaURL == nil {
		return errors.New("no previous location areas")
	}
	response, error := config.pokeapiClient.ListLocationAreas(config.previousLocationAreaURL)
	if error != nil {
		return error
	}
	fmt.Println("")
	fmt.Println("Location areas:")
	for _, area := range response.Results {
		fmt.Printf("%s\n", area.Name)
	}
	fmt.Println("")
	config.nextLocationAreaURL = response.Next
	config.previousLocationAreaURL = response.Previous
	return nil
}