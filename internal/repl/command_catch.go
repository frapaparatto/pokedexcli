package repl

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
)

const catchConstant = 100.0

func catchChance(baseExp int) float64 {
	p := catchConstant / (catchConstant + float64(baseExp))
	return math.Min(math.Max(p, 0.05), 0.90)
}

func attemptCatch(baseExp int) bool {
	return rand.Float64() < catchChance(baseExp)
}

func commandCatch(conf *Config, options ...string) error {
	if len(options) == 0 {
		return errors.New("usage: catch <pokemon-name>")
	}

	pokemon := options[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonResp, err := conf.pokeapiClient.GetPokemon(pokemon)

	if err != nil {
		return err
	}

	catched := attemptCatch(pokemonResp.BaseExperience)
	if !catched {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		conf.pokedex.Add(pokemonResp)
	}

	return nil
}
