package main

import (
	"fmt"
	"log"
)

func callbackMap(config *config) error {
	response, error := config.pokeapiClient.ListLocationAreas(config.nextLocationAreaURL)
	if error != nil {
		log.Fatal(error)
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