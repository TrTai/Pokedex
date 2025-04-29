package main

import (
	"fmt"
	//"net/http"
	"os"
	"time"

	"github.com/TrTai/pokeapi"
	"github.com/TrTai/pokecache"
)

var cmdMap map[string]pokeapi.CliCommand
var conf pokeapi.Config
var pokeCache *pokecache.Cache

func init() {
	cmdMap = map[string]pokeapi.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays location list",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays last location list",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays Pokemon at location",
			Callback:    commandExplore,
		},
	}
	pokeCache = pokecache.NewCache(30 * time.Minute)
}

func commandExit(c *pokeapi.Config, txt []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *pokeapi.Config, txt []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmdMap {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func commandMap(c *pokeapi.Config, txt []string) error {
	if c.NextURL == "" {
		c.NextURL = "https://pokeapi.co/api/v2/location-area/"
	}
	err := pokeapi.PokeGetLocs(c, c.NextURL, pokeCache)
	if err != nil {
		return fmt.Errorf("failed to get locations: %w", err)
	}
	return nil
}

func commandMapb(c *pokeapi.Config, txt []string) error {
	if c.PreviousURL == "" || c.PreviousURL == "https://pokeapi.co/api/v2/location-area/" {
		fmt.Println("This is the first page.")
		c.PreviousURL = "https://pokeapi.co/api/v2/location-area/"
	}
	err := pokeapi.PokeGetLocs(c, c.PreviousURL, pokeCache)
	if err != nil {
		return fmt.Errorf("failed to get locations: %w", err)
	}
	return nil
}

func commandExplore(c *pokeapi.Config, txt []string) error {
	if txt == nil {
		fmt.Println("Please provide a location name.")
		return nil
	} else if len(txt) > 1 {
		fmt.Println("Please provide one location area name.")
		return nil
	}
	locName := txt[0]
	locURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locName)
	err := pokeapi.PokeGetExplore(c, locURL, pokeCache)
	if err != nil {
		return fmt.Errorf("failed to get encounters: %w", err)
	}
	return nil
}
