package pokedex

import "testing"

func TestNewPokedexIsEmpty(t *testing.T) {
	p := NewPokedex()

	if p.pokemons == nil {
		t.Fatalf("expected pokemons map to be initialized, got nil")
	}

	if len(p.Names()) != 0 {
		t.Errorf("expected a new pokedex to be empty, got %d entries", len(p.Names()))
	}
}

func TestPokedexAdd(t *testing.T) {
	p := NewPokedex()
	pikachu := Pokemon{Name: "pikachu"}

	p.Add(pikachu)

	got, ok := p.Get("pikachu")
	if !ok {
		t.Fatalf("expected pikachu to be in the pokedex after Add")
	}

	if got.Name != "pikachu" {
		t.Errorf("got %v, expected %v", got.Name, "pikachu")
	}

	if len(p.Names()) != 1 {
		t.Errorf("got %d entries, expected 1", len(p.Names()))
	}
}

func TestPokedexAddOverwritesExisting(t *testing.T) {
	p := NewPokedex()
	p.Add(Pokemon{Name: "pikachu", BaseExperience: 1})
	p.Add(Pokemon{Name: "pikachu", BaseExperience: 2})

	if len(p.Names()) != 1 {
		t.Fatalf("got %d entries, expected 1 (re-catch should overwrite)", len(p.Names()))
	}

	got, _ := p.Get("pikachu")
	if got.BaseExperience != 2 {
		t.Errorf("got BaseExperience %d, expected 2", got.BaseExperience)
	}
}

func TestPokedexGetNotFound(t *testing.T) {
	p := NewPokedex()

	if _, ok := p.Get("pikachu"); ok {
		t.Errorf("expected ok=false for a pokemon that was never caught")
	}
}
