package commands

import (
	"fmt"
	"os"

	"github.com/zopitty/pokedexcli/config"
)

func commandExit(cfg *config.Config, args ...string) error {
    fmt.Println("Exiting the Pokedex...")
    os.Exit(0)
    return nil
}
