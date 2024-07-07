package config

import "github.com/zopitty/pokedexcli/internal/pokeapi"

// shared state of the application
type Config struct {
	PokeapiClient   pokeapi.Client
	NextLocationURL *string
	PrevLocationURL *string
	CaughtPokemon   map[string]pokeapi.Pokemon
}
