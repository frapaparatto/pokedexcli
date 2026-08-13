package repl

import (
	"errors"
	"fmt"
)

// commandExplore prints all Pokemon found in the given location area.
// Usage: explore <location-area-name>
func commandExplore(conf *Config, options ...string) error {
	if len(options) == 0 {
		return errors.New("usage: explore <location-area>")
	}

	location := options[0]
	fmt.Printf("Exploring %s...\n", location)

	locationResp, err := conf.pokeapiClient.ListPokemon(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
