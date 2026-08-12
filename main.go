package main

import "github.com/frapaparatto/pokedexcli/internal/repl"

func main() {
	conf := repl.NewConfig()
	conf.Next = "https://pokeapi.co/api/v2/location-area/"
	repl.Start(conf)
}
