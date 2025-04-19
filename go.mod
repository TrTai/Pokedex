module github.com/TrTai/pokedexcli

go 1.24.1

replace github.com/TrTai/pokecache => ./internal/pokecache

replace github.com/TrTai/pokeapi => ./internal/pokeapi

require (
	github.com/TrTai/pokeapi v0.0.0-00010101000000-000000000000
	github.com/TrTai/pokecache v0.0.0-00010101000000-000000000000
)
