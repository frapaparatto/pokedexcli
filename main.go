package main

import (
	"strings"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

var commands map[string]cliCommands

func main() {
	// TODO: before moving on, refactor and add eventual tests
	commands = map[string]cliCommands{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	startRepl()

}
