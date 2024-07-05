package commands

import (
	"errors"
	"fmt"

	"github.com/zopitty/pokedexcli/config"
)

func commandMapB(cfg *config.Config, args ...string) error {
    if cfg.PrevLocationURL == nil {
        return errors.New("this is the first page")
    }
    fmt.Println()
    fmt.Println("=================================")
    fmt.Println("Displaying Previous PokeLocations")
    fmt.Println("=================================")
    res, err := cfg.PokeapiClient.ListLocationAreas(cfg.PrevLocationURL)
    if err != nil {
        return err
    }
    for _, area := range res.Results{
        fmt.Printf(" - %s\n", area.Name)
    }
    cfg.NextLocationURL = res.Next
    cfg.PrevLocationURL = res.Previous
    return nil
}
