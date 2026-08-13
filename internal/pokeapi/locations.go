package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations fetches a page of location areas. When pageURL is nil, the
// first page is requested; otherwise pageURL (a Next/Previous link from a
// prior response) is used. Responses are served from cache when available.
func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		var locationsResp RespLocations
		if err := json.Unmarshal(val, &locationsResp); err != nil {
			return RespLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	var locationsResp RespLocations
	if err := json.Unmarshal(dat, &locationsResp); err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

// ListPokemon fetches a single location area and returns the Pokemon found
// there. Responses are served from cache when available.
func (c *Client) ListPokemon(location string) (RespLocation, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		var locationResp RespLocation
		if err := json.Unmarshal(val, &locationResp); err != nil {
			return RespLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return RespLocation{}, fmt.Errorf("location %q not found", location)
	}
	if resp.StatusCode != http.StatusOK {
		return RespLocation{}, fmt.Errorf("unexpected status code %d fetching location %q", resp.StatusCode, location)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	var locationResp RespLocation
	if err := json.Unmarshal(dat, &locationResp); err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
