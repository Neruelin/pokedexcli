package pokeapi

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
	"github.com/neruelin/pokedexcli/internal/pokecache"
)

const baseURL string = "https://pokeapi.co/api"

type JsonGetLocations struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

type JsonGetLocation struct {
	Id int `json:"id"`
	Name string `json:"name"`
	GameIndex int `json:"game_index"`
	EncounterMethodRates []struct{
		EncounterMethod struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct{
			Rate int `json:"rate"`
			Version struct{
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"location"`
	Names []struct{
		Name string `json:"name"`
		Language struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"language"`
	}
	PokemonEncounters []struct{
		Pokemon struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct{
			MaxChance int `json:"max_chance"`
			Version struct{
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version"`
			EncounterDetails []struct{
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				ConditionValues []struct{
					Id int `json:"id"`
					Name string `json:"name"`
					Condition struct{
						Name string `json:"name"`
						Url string `json:"url"`
					} `json:"condition"`
					Names []struct{
						Name string `json:"name"`
						Language struct{
							Name string `json:"name"`
							Url string `json:"url"`
						} `json:"language"`
					} `json:"names"`
				} `json:"condition_values"`
				Chance int `json:"chance"`
				Method struct{
					Name string `json:"name"`
					Url string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) getRequest(url string) (data []byte, err error) {
	// Check cache
	if cacheValue, ok := c.cache.Get(url); ok {
		data = cacheValue
	} else {
		// Get request
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, err	
		}
		// Read from body stream
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return []byte{}, fmt.Errorf("Response failed with status code %v", res.StatusCode)
		}
		if err != nil {
			return []byte{}, err
		}
		c.cache.Set(url, body)
		data = body
	}
	return 
}

func (c *Client) GetLocations(url string) ([]string, string, string, error) {

	if url == "" {
		url = baseURL + "/v2/location-area"
	}

	data, err := c.getRequest(url)
	
	// Converting JSON string to struct
	jsonLocations := JsonGetLocations{}
	err = json.Unmarshal(data, &jsonLocations)
	if err != nil {
		return []string{}, "", "", err
	}

	locationList := make([]string, len(jsonLocations.Results))

	for i, location := range jsonLocations.Results {
		locationList[i] = location.Name
	}

	return locationList, jsonLocations.Previous, jsonLocations.Next, nil
}

func (c *Client) GetLocation(name string) ([]string, error) {
	data, err := c.getRequest(baseURL + "/v2/location-area/" + name)
	if err != nil {
		return []string{}, err
	}

	// Converting JSON string to struct
	jsonGetLocation := JsonGetLocation{}
	err = json.Unmarshal(data, &jsonGetLocation)
	if err != nil {
		return []string{}, err
	}

	pokemonList := make([]string, len(jsonGetLocation.PokemonEncounters))

	for idx, pokemonEncounter := range jsonGetLocation.PokemonEncounters {
		pokemonList[idx] = pokemonEncounter.Pokemon.Name
	}

	return pokemonList, nil
}

type Client struct {
	cache pokecache.PokeCache
}

func NewClient() Client {
	return Client{cache: pokecache.New()}
}