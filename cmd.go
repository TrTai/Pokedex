package main

import (
	"fmt"
	//"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL     string
	previousURL string
}

var cmdMap map[string]cliCommand
var conf config

func init() {
	cmdMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location list",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays last location list",
			callback:    commandMapb,
		},
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *config) error {
	if c.nextURL == "" {
		c.nextURL = "https://pokeapi.co/api/v2/location-area/"
	}
	err := pokeGetLocs(c, c.nextURL)
	if err != nil {
		return fmt.Errorf("failed to get locations: %w", err)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.previousURL == "" || c.previousURL == "https://pokeapi.co/api/v2/location-area/" {
		fmt.Println("This is the first page.")
		c.previousURL = "https://pokeapi.co/api/v2/location-area/"
	}
	err := pokeGetLocs(c, c.previousURL)
	if err != nil {
		return fmt.Errorf("failed to get locations: %w", err)
	}
	return nil
}
