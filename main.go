package main

import (
	"time"

	"github.com/frapaparatto/pokedexcli/internal/pokecache"
	"github.com/frapaparatto/pokedexcli/internal/repl"
)

func main() {
	cache := pokecache.NewCache(5 * time.Millisecond)

	conf := repl.NewConfig()
	conf.Next = "https://pokeapi.co/api/v2/location-area/"
	repl.Start(conf)
}
