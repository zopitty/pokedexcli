package main

import (
	"fmt"
	"time"

	"github.com/zopitty/pokedexcli/config"
	"github.com/zopitty/pokedexcli/internal/pokeapi"
	"github.com/zopitty/pokedexcli/repl"
)

func main() {
	cfg := config.Config{
		PokeapiClient: pokeapi.NewClient(time.Hour),
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	// fmt.Println("WELCOME TO THE POKEDEX!!")
	fmt.Println("Welcome to the Pokedex!")
	repl.StartREPL(&cfg)
}
