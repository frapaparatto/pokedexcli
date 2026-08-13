package repl

import (
	"testing"

	"github.com/frapaparatto/pokedexcli/internal/pokeapi"
	"github.com/frapaparatto/pokedexcli/internal/pokedex"
)

func TestCommandInspectNotCaught(t *testing.T) {
	conf := &Config{pokedex: pokedex.NewPokedex()}

	err := commandInspect(conf, "pidgey")
	if err == nil {
		t.Fatalf("expected an error for an uncaught pokemon, got nil")
	}
}

func TestCommandInspectNoArgs(t *testing.T) {
	conf := &Config{pokedex: pokedex.NewPokedex()}

	if err := commandInspect(conf); err == nil {
		t.Fatalf("expected a usage error when no name is given, got nil")
	}
}

func TestCommandInspectCaught(t *testing.T) {
	conf := &Config{pokedex: pokedex.NewPokedex()}
	conf.pokedex.Add(pokeapi.RespPokemon{
		Name:   "pidgey",
		Height: 3,
		Weight: 18,
	})

	if err := commandInspect(conf, "pidgey"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
