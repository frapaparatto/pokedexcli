package repl

import (
	"fmt"
	"os"
)

// TODO: don't change messages until the project is completed, then refactor
func commandExit(conf *Config, options ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
