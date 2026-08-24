package repl

import (
	"testing"

	"github.com/frapaparatto/pokedexcli/internal/pokedex"
)

func TestCommandPokedexEmpty(t *testing.T) {
	conf := &Config{pokedex: pokedex.NewPokedex()}

	if err := commandPokedex(conf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCommandPokedexWithCaughtPokemon(t *testing.T) {
	conf := &Config{pokedex: pokedex.NewPokedex()}
	conf.pokedex.Add(pokedex.Pokemon{Name: "pidgey"})
	conf.pokedex.Add(pokedex.Pokemon{Name: "pikachu"})

	if err := commandPokedex(conf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
