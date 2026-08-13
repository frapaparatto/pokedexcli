package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

	if val, ok := c.cache.Get(url); ok {
		var pokemonResp RespPokemon
		if err := json.Unmarshal(val, &pokemonResp); err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return RespPokemon{}, fmt.Errorf("pokemon %q not found", name)
	}
	if resp.StatusCode != http.StatusOK {
		return RespPokemon{}, fmt.Errorf("unexpected status code %d fetching pokemon %q", resp.StatusCode, name)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	var pokemonResp RespPokemon
	if err := json.Unmarshal(dat, &pokemonResp); err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
