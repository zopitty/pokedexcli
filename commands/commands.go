package commands

import "github.com/zopitty/pokedexcli/config"

type CliCommand struct {
    Name        string
    Description string
    Callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
    return map[string]CliCommand {
        "help": {
            Name:        "help",
            Description: "Displays a help message",
            Callback:    commandHelp,
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
        "explore": {
            Name:        "explore {location area}",
            Description: "show pokemon in the area",
            Callback:    commandExplore,
        },
        "catch": {
            Name:        "catch {pokemon name}",
            Description: "catch pokemon",
            Callback:    commandCatch,
        },
        "exit": {
            Name:        "exit",
            Description: "Exit the Pokedex",
            Callback:    commandExit,
        },
    }
}
