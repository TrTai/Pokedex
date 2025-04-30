package pokeapi

import (
	//"encoding/json"
	"fmt"
	"github.com/TrTai/pokecache"
	//"net/http"
)

type exploreResponse struct {
	_                  any                `json:"encounter_method_rates"`
	_                  any                `json:"location"`
	_                  any                `json:"names"`
	pokemon_encounters []pokemonencounter `json:"pokemon_encounters"`
}

type pokemonencounter struct {
	Pokemon pokemon `json:"pokemon"`
	_       any     `json:"version_details"`
}
type pokemon struct {
	Name string `json:"name"`
	_    string `json:"url"`
}

func PokeGetExplore(conf *Config, url string, pc *pokecache.Cache) error {
	exploreResponse, err := pokeGetCached(conf, url, pc)
	if err != nil {
		return fmt.Errorf("failed to get explore data: %w", err)
	}
	encounterList := exploreResponse.pokemon_encounters
	fmt.Println("Pokemon Encounters:")
	fmt.Println(encounterList)
	for i, encounter := range encounterList {
		fmt.Println(i)
		fmt.Printf("Pokemon: %s\n", encounter.Pokemon.Name)
	}
	return nil
}
