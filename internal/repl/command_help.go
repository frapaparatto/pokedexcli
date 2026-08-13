package repl

import (
	"fmt"
	"sort"
)

func commandHelp(conf *Config, options ...string) error {
	fmt.Printf("Usage:\n\n\t<command> [arguments]\n")
	fmt.Printf("\nCommands:\n\n")

	names := make([]string, 0, len(conf.commands))
	longest := 0
	for name := range conf.commands {
		names = append(names, name)
		if len(name) > longest {
			longest = len(name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		cmd := conf.commands[name]
		fmt.Printf("\t%-*v\t\t%v\n", longest, cmd.name, cmd.description)
	}

	fmt.Printf("\n")
	return nil
}
