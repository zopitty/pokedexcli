package commands

import (
	"fmt"

	"github.com/zopitty/pokedexcli/config"
)

func commandMap(cfg *config.Config) error {

    fmt.Println()
    fmt.Println("========================")
    fmt.Println("Displaying PokeLocations")
    fmt.Println("========================")
    res, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationURL)
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
