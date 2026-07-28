package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
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
		c, ok := commands[command]

		if !ok {
			fmt.Printf("Unknown command\n")
		} else {
			c.callback()
		}
	}
}
