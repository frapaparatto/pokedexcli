package main

import (
	"time"

	"github.com/frapaparatto/pokedexcli/internal/pokeapi"
	"github.com/frapaparatto/pokedexcli/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := repl.NewConfig(pokeClient)
	repl.Start(conf)
}
