package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func main() {
	cmdMap := map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}

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

		if c, ok := cmdMap[cmd]; ok {
			err := c.callback()
			if err != nil {
				fmt.Println("%v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

