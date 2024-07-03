package main

import (
	"fmt"

	"github.com/zopitty/pokedexcli/config"
	"github.com/zopitty/pokedexcli/internal/pokeapi"
	"github.com/zopitty/pokedexcli/repl"
)


func main() {
    cfg := config.Config{
        PokeapiClient: pokeapi.NewClient(),

    }
    fmt.Println("WELCOME TO THE POKEDEX!!")
    repl.StartREPL(&cfg)
}
