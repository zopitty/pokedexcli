package commands

import (
	"fmt"

	"github.com/zopitty/pokedexcli/config"
)

func commandPokedex(cfg *config.Config, args ...string) error {
    fmt.Println("Pokemon in Pokedex:")

    for _, caughtPokemon := range cfg.CaughtPokemon {
        fmt.Printf(" - %s\n", caughtPokemon.Name)
    }

    return nil 
}
