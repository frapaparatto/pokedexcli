package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// Start runs the REPL loop, reading commands from stdin until the process exits.
func Start(conf *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		var options []string

		if len(cleaned) > 1 {
			options = cleaned[1:]
		}

		c, ok := conf.commands[command]

		if !ok {
			fmt.Printf("Unknown command\n")
		} else if err := c.callback(conf, options...); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
