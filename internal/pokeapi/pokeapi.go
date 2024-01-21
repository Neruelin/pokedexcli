package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"github.com/neruelin/pokedexcli/internal/pokecache"
)

type Location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type JsonLocations struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}

func (c *Client) GetLocations(url string) ([]string, string, string, error) {

	var data []byte

	// Check cache
	if cacheValue, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit ", url)
		data = cacheValue
	} else {
		fmt.Println("Cache miss ", url)
		// Get request
		res, err := http.Get(url)
		if err != nil {
			return []string{}, "", "", err	
		}
		// Read from body stream
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return []string{}, "", "", fmt.Errorf("Response failed with status code %v", res.StatusCode)
		}
		if err != nil {
			return []string{}, "", "", err
		}
		c.cache.Set(url, body)
		data = body
	}
	
	// Converting JSON string to struct
	jsonLocations := JsonLocations{}
	err := json.Unmarshal(data, &jsonLocations)
	if err != nil {
		return []string{}, "", "", err
	}

	locationList := make([]string, len(jsonLocations.Results))

	for i, location := range jsonLocations.Results {
		locationList[i] = location.Name
	}

	return locationList, jsonLocations.Previous, jsonLocations.Next, nil
}

type Client struct {
	cache pokecache.PokeCache
}

func NewClient() Client {
	return Client{cache: pokecache.New()}
}