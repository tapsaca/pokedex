package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println(("Invalid command"))
			continue
		}
		error := command.callback(config, args...)
		if error != nil {
			fmt.Println(error)
		}
	}
}

func cleanInput(str string) []string {
	return strings.Fields(strings.ToLower(str))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"catch": {
			name: "catch {pokemon}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback: callbackCatch,
		},
		"exit": {
			name: "exit",
			description: "closes the program",
			callback: callbackExit,
		},
		"explore": {
			name: "explore {location-area}",
			description: "Lists the pokemon in a location area",
			callback: callbackExplore,
		},
		"help": {
			name: "help",
			description: "prints the help menu",
			callback: callbackHelp,
		},
		"map": {
			name: "map",
			description: "prints next location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "prints previous location areas",
			callback: callbackMapb,
		},
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}