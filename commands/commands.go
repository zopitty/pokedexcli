package commands

import "github.com/zopitty/pokedexcli/config"

type CliCommand struct {
    Name        string
    Description string
    Callback    func(*config.Config) error
}

func GetCommands() map[string]CliCommand {
    return map[string]CliCommand {
        "help": {
            Name:        "help",
            Description: "Displays a help message",
            Callback:    commandHelp,
        },
        "exit": {
            Name:        "exit",
            Description: "Exit the Pokedex",
            Callback:    commandExit,
        },
        "map": {
            Name:        "map",
            Description: "Display 20 areas",
            Callback:    commandMap,
        },
        "mapb": {
            Name:        "mapb",
            Description: "Display previous 20 areas",
            Callback:    commandMapB,
        },
    }
}
