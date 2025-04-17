package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func pokeGetLocs(c *config) error {
	resp, err := http.Get(c.nextURL)
	if err != nil {
		return fmt.Errorf("failed to fetch location data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("PokeAPI Http statusCode not-OK, StatusCode: %s", resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	var locationsResponse mapResponse
	dErr := decoder.Decode(&locationsResponse)
	if dErr != nil {
		return fmt.Errorf("failed to decode location data: %w", dErr)
	}
	locations := locationsResponse.Results
	for _, loc := range locations {
		fmt.Printf("Location: %s\n", loc.Name)
	}
	if locationsResponse.Next != "" {
		c.nextURL = locationsResponse.Next
	}
	if locationsResponse.Previous != "" {
		c.previousURL = locationsResponse.Previous
	} else {
		c.previousURL = ""
	}
	return nil
}
