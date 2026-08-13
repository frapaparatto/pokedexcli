package repl

import (
	"errors"
	"fmt"
)

// commandInspect prints details about a Pokemon the user has already caught.
// Usage: inspect <pokemon-name>
func commandInspect(conf *Config, options ...string) error {
	if len(options) == 0 {
		return errors.New("usage: inspect <pokemon-name>")
	}

	name := options[0]
	pokemon, ok := conf.pokedex.Pokemons[name]
	if !ok {
		return fmt.Errorf("you have not caught %s", name)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
