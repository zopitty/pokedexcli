package repl

import (
    "bufio"
    "fmt"
    "os"
    "github.com/zopitty/pokedexcli/commands"

)

func StartREPL() {
    scanner := bufio.NewScanner(os.Stdin)
    cmds := commands.GetCommands()
    for {
        fmt.Print("Pokedex > ")
        if !scanner.Scan() {
            break
        }

        input := scanner.Text()
        if cmd, exists := cmds[input]; exists {
            if err := cmd.Callback(); err != nil {
                fmt.Println("Error:", err)
            }
        } else {
            fmt.Println("Unknown command")
        }
    }
}
