package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zopitty/pokedexcli/commands"
	"github.com/zopitty/pokedexcli/config"
)

func StartREPL(cfg *config.Config) {
    scanner := bufio.NewScanner(os.Stdin)
    cmds := commands.GetCommands()
    for {
        fmt.Print("Pokedex > ")
        if !scanner.Scan() {
            break
        }

        input := scanner.Text()

        parts := strings.Fields(input)

        if len(parts) == 0 {
            continue
        }
        args := []string{}
        command := parts[0]
        if len(parts) > 1 {
            args = parts[1:]
        }
        if cmd, exists := cmds[command]; exists {
            if err := cmd.Callback(cfg, args...); err != nil {
                fmt.Println("Error:", err)
            }
        } else {
            fmt.Println("Unknown command")
        }
    }
}

