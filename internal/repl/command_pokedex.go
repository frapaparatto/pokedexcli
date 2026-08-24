package repl

import (
	"fmt"
	"sort"
)

// commandPokedex prints the names of every Pokemon the user has caught so far.
func commandPokedex(conf *Config, options ...string) error {
	names := conf.pokedex.Names()
	sort.Strings(names)

	fmt.Println("Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
