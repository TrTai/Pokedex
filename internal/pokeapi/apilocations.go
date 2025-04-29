package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/TrTai/pokecache"
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

func PokeGetLocs(conf *Config, url string, pc *pokecache.Cache) error {

	locationsResponse, err := pokeGetLocsCached(conf, url, pc)
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

func pokeGetLocsCached(conf *Config, url string, pc *pokecache.Cache) (mapResponse, error) {
	cachedData, ok := pc.Get(url)
	if ok {
		var locations mapResponse
		mErr := json.Unmarshal(cachedData, &locations)
		if mErr != nil {
			return locations, fmt.Errorf("failed to unmarshal cached data: %w", mErr)
		}
		fmt.Println("Using cached data")
		return locations, nil
	} else {
		locations, err := pokeGetLocsHttp(conf, url, pc)
		if err != nil {
			return locations, fmt.Errorf("failed to get locations from HTTP: %w", err)
		}
		return locations, nil
	}

}

func pokeGetLocsHttp(conf *Config, url string, pc *pokecache.Cache) (mapResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return mapResponse{}, fmt.Errorf("failed to fetch location data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return mapResponse{}, fmt.Errorf("PokeAPI Http statusCode not-OK, StatusCode: %s", resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	var locationsResponse mapResponse
	dErr := decoder.Decode(&locationsResponse)
	if dErr != nil {
		return mapResponse{}, fmt.Errorf("failed to decode location data: %w", dErr)
	}
	fmt.Println("Fetched data from PokeAPI")
	if err != nil {
		return mapResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	locationBytes, err := json.Marshal(locationsResponse)
	pc.Add(url, locationBytes)

	return locationsResponse, nil
}
