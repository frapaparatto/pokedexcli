package pokedex

import (
	"testing"

	"github.com/frapaparatto/pokedexcli/internal/pokeapi"
)

func TestNewPokedexIsEmpty(t *testing.T) {
	p := NewPokedex()

	if p.Pokemons == nil {
		t.Fatalf("expected Pokemons map to be initialized, got nil")
	}

	if len(p.Pokemons) != 0 {
		t.Errorf("expected a new pokedex to be empty, got %d entries", len(p.Pokemons))
	}
}

func TestPokedexAdd(t *testing.T) {
	p := NewPokedex()
	pikachu := pokeapi.RespPokemon{Name: "pikachu"}

	p.Add(pikachu)

	got, ok := p.Pokemons["pikachu"]
	if !ok {
		t.Fatalf("expected pikachu to be in the pokedex after Add")
	}

	if got.Name != "pikachu" {
		t.Errorf("got %v, expected %v", got.Name, "pikachu")
	}

	if len(p.Pokemons) != 1 {
		t.Errorf("got %d entries, expected 1", len(p.Pokemons))
	}
}

func TestPokedexAddOverwritesExisting(t *testing.T) {
	p := NewPokedex()
	p.Add(pokeapi.RespPokemon{Name: "pikachu", BaseExperience: 1})
	p.Add(pokeapi.RespPokemon{Name: "pikachu", BaseExperience: 2})

	if len(p.Pokemons) != 1 {
		t.Fatalf("got %d entries, expected 1 (re-catch should overwrite)", len(p.Pokemons))
	}

	if p.Pokemons["pikachu"].BaseExperience != 2 {
		t.Errorf("got BaseExperience %d, expected 2", p.Pokemons["pikachu"].BaseExperience)
	}
}
