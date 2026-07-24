package main

import (
	"bufio"
	"fmt"
	"os"
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

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		command := cleanInput(input)[0]
		c, ok := commands[command]

		if !ok {
			fmt.Printf("Unknown command\n")
		} else {
			c.callback()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
