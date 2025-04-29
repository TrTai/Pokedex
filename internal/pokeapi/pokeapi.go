package pokeapi

/* import (
	"encoding/json"
	"fmt"
	"github.com/TrTai/pokecache"
	"net/http"
) */

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

type Config struct {
	NextURL     string
	PreviousURL string
}
