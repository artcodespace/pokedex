package main

type config struct {
	Next *string
	Previous *string
}

type cliCommand struct {
	name string
	description string
	callback func(config *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
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
		"map" : {
			name: "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb" : {
			name: "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback: commandMapb,
		},
	}
}
