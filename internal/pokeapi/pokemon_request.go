package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint
	// Check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		pokemon := Pokemon{}
		error := json.Unmarshal(data, &pokemon)
		if error != nil {
			return Pokemon{}, error
		}
		return pokemon, nil
	}
	request, error := http.NewRequest("GET", fullURL, nil)
	if error != nil {
		return Pokemon{}, error
	}
	response, error := c.httpClient.Do(request)
	if error != nil {
		return Pokemon{}, error
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("status code: %v", response.StatusCode)
	}
	data, error = io.ReadAll(response.Body)
	if error != nil {
		return Pokemon{}, error
	}
	pokemon := Pokemon{}
	error = json.Unmarshal(data, &pokemon)
	if error != nil {
		return Pokemon{}, error
	}
	c.cache.Add(fullURL, data)
	return pokemon, nil
}