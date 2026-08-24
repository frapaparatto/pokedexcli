package repl

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/frapaparatto/pokedexcli/internal/pokeapi"
	"github.com/frapaparatto/pokedexcli/internal/pokedex"
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

	caught := attemptCatch(pokemonResp.BaseExperience)
	if !caught {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		conf.pokedex.Add(toPokedexPokemon(pokemonResp))
	}

	return nil
}

// toPokedexPokemon converts a PokeAPI response into the pokedex package's
// own Pokemon type, keeping the domain package free of any dependency on
// pokeapi.
func toPokedexPokemon(resp pokeapi.RespPokemon) pokedex.Pokemon {
	stats := make([]pokedex.PokemonStat, 0, len(resp.Stats))
	for _, s := range resp.Stats {
		stats = append(stats, pokedex.PokemonStat{Name: s.Stat.Name, Value: s.BaseStat})
	}

	types := make([]string, 0, len(resp.Types))
	for _, t := range resp.Types {
		types = append(types, t.Type.Name)
	}

	return pokedex.Pokemon{
		Name:           resp.Name,
		BaseExperience: resp.BaseExperience,
		Height:         resp.Height,
		Weight:         resp.Weight,
		Stats:          stats,
		Types:          types,
	}
}
