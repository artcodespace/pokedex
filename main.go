package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		cleanedInput := cleanInput(line)

		if len(cleanedInput) == 0 {
			fmt.Println("You did not enter any input!")
		} else {
			fmt.Println("Your command was:", cleanInput(line)[0])
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

