package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		cleanedLine := cleanInput(line)

		if len(cleanedLine) == 0 {
			fmt.Println("No command entered")
			continue
		}

		cmd := cleanedLine[0]

		if c, ok := getCommands()[cmd]; ok {
			config := config {
				Next: "",
				Previous: "",
			}
			err := c.callback(&config)
			if err != nil {
				fmt.Printf("%v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
