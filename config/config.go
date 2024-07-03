package config

import "github.com/zopitty/pokedexcli/internal/pokeapi"

type Config struct {
    PokeapiClient   pokeapi.Client
    NextLocationURL *string
    PrevLocationURL *string
}
