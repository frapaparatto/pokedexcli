package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Config holds shared REPL state, including the registered commands.
type Config struct {
	commands map[string]cliCommand
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

// NewConfig builds a Config with all CLI commands registered.
func NewConfig() *Config {
	conf := &Config{}
	conf.commands = map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapb,
		},
	}
	return conf
}

// TODO: don't change messages until the project is completed, then refactor
func commandExit(conf *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range conf.commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}

type Result struct {
	Count     int    `json:"count"`
	Next      string `json:"next"`
	Previous  string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(conf *Config) error {
	// TODO: understand what *string means
	res, err := http.Get(conf.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	for _, location := range result.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	conf.Next = result.Next
	conf.Previous = result.Previous

	return nil
}

func commandMapb(conf *Config) error {
	res, err := http.Get(conf.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	for _, location := range result.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	conf.Next = result.Next
	conf.Previous = result.Previous

	return nil
}
