package repl

import "fmt"

func commandHelp(conf *Config, options ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range conf.commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}
