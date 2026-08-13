package pokedex

import "github.com/frapaparatto/pokedexcli/internal/pokeapi"

// Pokedex keeps track of every pokemon that has been caught, keyed by name.
type Pokedex struct {
	Pokemons map[string]pokeapi.RespPokemon
}

func (p *Pokedex) Add(pokemon pokeapi.RespPokemon) {
	p.Pokemons[pokemon.Name] = pokemon
}

// NewPokedex creates an empty Pokedex.
func NewPokedex() *Pokedex {
	return &Pokedex{
		Pokemons: make(map[string]pokeapi.RespPokemon),
	}
}
