package commands

import "fmt"

func commandHelp() error {
    fmt.Println("=======================")
    fmt.Println("Welcome to the pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range GetCommands() {
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }
    fmt.Println()
    fmt.Println("=======================")
    return nil
}
