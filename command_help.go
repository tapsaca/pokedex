package main

import "fmt"

func callbackHelp(config *config) error {
	fmt.Println("")
	fmt.Println("Available commands:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}