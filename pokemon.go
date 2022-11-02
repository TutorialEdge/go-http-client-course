package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetPokemonByName - note
func (c *Client) GetPokemonByName(ctx context.Context, name string) (Pokemon, error) {
	// build the request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://pokeapi.co/api/v2/pokemon/"+name,
		nil,
	)
	if err != nil {
		return Pokemon{}, err
	}

	// Add request headers - in this case we want to specify
	// that we accept JSON. You could also add any o11y headers
	// such as trace headers or request IDs.
	req.Header.Add("Accept", "application/json")

	// actually perform the http request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	// ALWAYS close the request body - otherwise you end up
	// in a world of pain
	defer resp.Body.Close()

	// Let's unmarshal the response body into a struct
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// Finally, we return the pokemon
	return pokemon, nil
}
