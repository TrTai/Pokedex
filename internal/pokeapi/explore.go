package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/TrTai/pokecache"
	"net/http"
)

type exploreResponse struct {
	pokemon_encounters []pokemonencounter `json:"pokemon_encounters"`
}

type pokemonencounter struct {
	Pokemon pokemon `json:"pokemon"`
}
type pokemon struct {
	Name string `json:"name"`
}

func PokeGetExplore(conf *Config, url string, pc *pokecache.Cache) error {
	exploreResponse, err := pokeGetExploreCached(conf, url, pc)
	if err != nil {
		return fmt.Errorf("failed to get explore data: %w", err)
	}
	encounterList := exploreResponse.pokemon_encounters
	fmt.Println("Pokemon Encounters:")
	for _, encounter := range encounterList {
		fmt.Printf("Pokemon: %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func pokeGetExploreCached(conf *Config, url string, pc *pokecache.Cache) (exploreResponse, error) {
	cachedData, ok := pc.Get(url)
	if ok {
		var exploreData exploreResponse
		mErr := json.Unmarshal(cachedData, &exploreData)
		if mErr != nil {
			return exploreData, fmt.Errorf("failed to unmarshal cached data: %w", mErr)
		}
		fmt.Println("Using cached data")
		return exploreData, nil
	} else {
		exploreData, err := pokeGetExploreHttp(conf, url, pc)
		if err != nil {
			return exploreData, fmt.Errorf("failed to get explore data from HTTP: %w", err)
		}
		return exploreData, nil
	}
}

func pokeGetExploreHttp(conf *Config, url string, pc *pokecache.Cache) (exploreResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return exploreResponse{}, fmt.Errorf("failed to get explore data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return exploreResponse{}, fmt.Errorf("failed to get explore data: %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	var exploreData exploreResponse
	dErr := decoder.Decode(&exploreData)
	if dErr != nil {
		return exploreResponse{}, fmt.Errorf("failed to decode explore data: %w", err)
	}

	exploreBytes, err := json.Marshal(exploreData)
	if err != nil {
		return exploreResponse{}, fmt.Errorf("failed to marshal explore data: %w", err)
	}
	pc.Add(url, exploreBytes)
	return exploreData, nil
}
