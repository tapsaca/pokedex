package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println(("Invalid command"))
			continue
		}
		command.callback()
	}
}

func cleanInput(str string) []string {
	return strings.Fields(strings.ToLower(str))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "closes the program",
			callback: callbackExit,
		},
		"help": {
			name: "help",
			description: "prints the help menu",
			callback: callbackHelp,
		},
	}
}

type cliCommand struct {
	name string
	description string
	callback func() error
}