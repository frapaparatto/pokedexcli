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

// ListLocations fetches a page of location areas. When pageURL is nil, the
// first page is requested; otherwise pageURL (a Next/Previous link from a
// prior response) is used. Responses are served from cache when available.
func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var resp RespLocations
	if err := c.getResource(url, &resp); err != nil {
		return RespLocations{}, err
	}
	return resp, nil
}

// ListPokemon fetches a single location area and returns the Pokemon found
// there. Responses are served from cache when available.
func (c *Client) ListPokemon(location string) (RespLocation, error) {
	url := baseURL + "/location-area/" + location

	var resp RespLocation
	if err := c.getResource(url, &resp); err != nil {
		return RespLocation{}, err
	}
	return resp, nil
}
