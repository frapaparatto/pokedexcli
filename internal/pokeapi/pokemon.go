package pokeapi

// RespPokemon is the PokeAPI response for a single pokemon resource
// (https://pokeapi.co/api/v2/pokemon/{name}), scoped to the fields the
// Pokedex cares about.
//
// Note: this endpoint doesn't include a plain "description" field. Flavor
// text lives on the separate "pokemon-species" resource
// (flavor_text_entries), so it isn't fetched here.
type RespPokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

// GetPokemon fetches a single pokemon by name. Responses are served from
// cache when available.
func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	var resp RespPokemon
	if err := c.getResource(url, &resp); err != nil {
		return RespPokemon{}, err
	}
	return resp, nil
}
