package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	// Check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResponse := LocationAreaResponse{}
		error := json.Unmarshal(data, &locationAreaResponse)
		if error != nil {
			return LocationAreaResponse{}, error
		}
		return locationAreaResponse, nil
	}
	request, error := http.NewRequest("GET", fullURL, nil)
	if error != nil {
		return LocationAreaResponse{}, error
	}
	response, error := c.httpClient.Do(request)
	if error != nil {
		return LocationAreaResponse{}, error
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("status code: %v", response.StatusCode)
	}
	data, error = io.ReadAll(response.Body)
	if error != nil {
		return LocationAreaResponse{}, error
	}
	locationAreaResponse := LocationAreaResponse{}
	error = json.Unmarshal(data, &locationAreaResponse)
	if error != nil {
		return LocationAreaResponse{}, error
	}
	c.cache.Add(fullURL, data)
	return locationAreaResponse, nil
}