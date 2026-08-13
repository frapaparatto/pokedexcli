package repl

import (
	"fmt"
	"sort"
)

// commandPokedex prints the names of every Pokemon the user has caught so far.
func commandPokedex(conf *Config, options ...string) error {
	names := make([]string, 0, len(conf.pokedex.Pokemons))
	for name := range conf.pokedex.Pokemons {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
