package pokedex

// Pokemon is the subset of a caught Pokemon's data the Pokedex keeps.
type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []PokemonStat
	Types          []string
}

// PokemonStat is a single named base stat (e.g. "speed": 90).
type PokemonStat struct {
	Name  string
	Value int
}

// Pokedex keeps track of every pokemon that has been caught, keyed by name.
type Pokedex struct {
	pokemons map[string]Pokemon
}

// NewPokedex creates an empty Pokedex.
func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemons: make(map[string]Pokemon),
	}
}

// Add records a caught pokemon, overwriting any existing entry with the same name.
func (p *Pokedex) Add(pokemon Pokemon) {
	p.pokemons[pokemon.Name] = pokemon
}

// Get returns the caught pokemon with the given name, and whether it was found.
func (p *Pokedex) Get(name string) (Pokemon, bool) {
	pokemon, ok := p.pokemons[name]
	return pokemon, ok
}

// Names returns the names of every caught pokemon, in no particular order.
func (p *Pokedex) Names() []string {
	names := make([]string, 0, len(p.pokemons))
	for name := range p.pokemons {
		names = append(names, name)
	}
	return names
}
