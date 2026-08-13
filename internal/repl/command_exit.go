package repl

import (
	"os"
)

func commandExit(conf *Config, options ...string) error {
	os.Exit(0)
	return nil
}
