package pokeapi

// RespLocations is the PokeAPI response for a page of location areas.
type RespLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// RespLocation is the PokeAPI response for a single location area, scoped to
// the fields needed to list the Pokemon encountered there.
type RespLocation struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
