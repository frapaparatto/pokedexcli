package pokeapi

import (
	"net/http"
	"time"

	"github.com/frapaparatto/pokedexcli/internal/pokecache"
)

// Client wraps an HTTP client and a response cache for talking to the PokeAPI.
type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

// NewClient builds a Client with the given request timeout and cache interval.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
