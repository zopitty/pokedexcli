package commands

import (
	"errors"
	"fmt"

	"github.com/zopitty/pokedexcli/config"
)

func commandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	caughtPokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

    fmt.Println("================")
    fmt.Println("Displaying Stats")
    fmt.Println("================")
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %v\n", caughtPokemon.Height)
	fmt.Printf("Weight: %v\n", caughtPokemon.Weight)
    fmt.Println("Stats:")
	for _, stat := range caughtPokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Type:")
	for _, typ := range caughtPokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
    fmt.Println("================")

	return nil
}
