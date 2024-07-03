package commands

import (
	"fmt"
	"os"

	"github.com/zopitty/pokedexcli/config"
)

func commandExit(cfg *config.Config) error {
    fmt.Println("Exiting the Pokedex...")
    os.Exit(0)
    return nil
}
