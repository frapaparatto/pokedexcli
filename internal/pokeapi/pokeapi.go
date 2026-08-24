package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

// getResource fetches url, serving from cache when available, and unmarshals
// the JSON response body into v (a pointer to the caller's concrete response
// type). On a cache hit, the cached bytes are unmarshaled directly. On a
// miss, the response is fetched, status-checked, unmarshaled, and cached for
// next time.
//
// getResource has no knowledge of what kind of resource it's fetching, so
// error messages are phrased in terms of url rather than a resource name.
func (c *Client) getResource(url string, v any) error {
	if val, ok := c.cache.Get(url); ok {
		return json.Unmarshal(val, v)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("%q not found", url)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d fetching %q", resp.StatusCode, url)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(dat, v); err != nil {
		return err
	}

	c.cache.Add(url, dat)
	return nil
}
