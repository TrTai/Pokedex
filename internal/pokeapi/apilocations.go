package pokeapi

import (
	//"encoding/json"
	"fmt"
	"github.com/TrTai/pokecache"
	//"net/http"
)

type mapResponse struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []mapResults `json:"results"`
}

type mapResults struct {
	Name   string `json:"name"`
	LocURL string `json:"url"`
}

func PokeGetLocs(conf *Config, url string, pc *pokecache.Cache) error {

	locationsResponse, err := pokeGetCached(conf, url, pc)
	if err != nil {
		return fmt.Errorf("failed to get locations: %w", err)
	}
	locations := locationsResponse.Results
	for _, loc := range locations {
		fmt.Printf("Location: %s\n", loc.Name)
	}
	if locationsResponse.Next != "" {
		conf.NextURL = locationsResponse.Next
	}
	if locationsResponse.Previous != "" {
		conf.PreviousURL = locationsResponse.Previous
	} else {
		conf.PreviousURL = ""
	}
	return nil
}
