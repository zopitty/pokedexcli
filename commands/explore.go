package commands

import (
	"errors"
	"fmt"

	"github.com/zopitty/pokedexcli/config"
)

func commandExplore(cfg *config.Config, args ...string) error {
    if len(args) != 1 {
        return errors.New("no location area provided")
    }

    locationAreaName := args[0]
    resLocationArea, err := cfg.PokeapiClient.GetLocationArea(locationAreaName)
    if err != nil {
        return err
    }
    fmt.Printf("pokemon in %s:\n", resLocationArea.Name)
    for _, pokemon := range resLocationArea.PokemonEncounters{
        fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
    }
    return nil
}
