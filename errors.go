package client

import "fmt"

type PokemonFetchErr struct {
	Message string
}

func (e PokemonFetchErr) Error() string {
	return fmt.Sprintf("failed to fetch a pokemon: %s", e.Message)
}
