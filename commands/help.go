package commands

import (
	"fmt"

	"github.com/zopitty/pokedexcli/config"
)

func commandHelp(cfg *config.Config, args ...string) error {
	var commandOrder = []string{
		"help",
		"map",
		"mapb",
		"explore",
		"catch",
		"exit",
	}
	cmds := GetCommands()
	fmt.Println("=======================")
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, key := range commandOrder {
		cmd := cmds[key]
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	
	}
	// for _, cmd := range GetCommands() {
	//     fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	// }
	fmt.Println()
	fmt.Println("=======================")
	return nil
}
