package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/zopitty/pokedexcli/config"
)

func commandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	resPokemon, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(resPokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	cfg.CaughtPokemon[pokemonName] = resPokemon
	fmt.Printf("%s was caught\n", pokemonName)
	return nil
}
