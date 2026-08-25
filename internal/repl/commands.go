package repl

import (
	"github.com/frapaparatto/pokedexcli/internal/pokeapi"
	"github.com/frapaparatto/pokedexcli/internal/pokedex"
)

// Config holds shared REPL state, including the registered commands.
type Config struct {
	commands         map[string]cliCommand
	pokeapiClient    *pokeapi.Client
	pokedex          *pokedex.Pokedex
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config, options ...string) error
}

// NewConfig builds a Config with all CLI commands registered.
func NewConfig(client *pokeapi.Client) *Config {
	return &Config{
		commands:      getCommands(),
		pokeapiClient: client,
		pokedex:       pokedex.NewPokedex(),
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List Pokémon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokémon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show details of a caught Pokémon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokémon",
			callback:    commandPokedex,
		},
	}
}
