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

type Pokemon struct {
	Abilities []struct{
		Ability struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot int `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"forms"` 
	GameIndices []struct{
		GameIndex int `json:"game_index"`
		Version struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"` 
	Height int `json:"height"`
	HeldItems []struct{
		Item struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"item"`
		VersionDetails []struct{
			Rarity int `json:"rarity"`
			Version struct{
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"` 
	Id int `json:"id"`
	IsDefault bool `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []struct{
		Move struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct{
			LevelLearnedAt int `json:"level_learned_at"`
			MoveLearnMethod struct{
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct{
				Name string `json:"name"`
				Url string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"` 
	Name string `json:"name"`
	Order int `json:"order"`
	Species struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"species"` 
	Sprites struct{
		BackDefault string `json:"back_default"`
		BackFemale string `json:"back_female"`
		BackShiny string `json:"back_shiny"`
		BackShinyFemale string `json:"back_shiny_female"`
		FrontDefault string `json:"front_default"`
		FrontFemale string `json:"front_female"`
		FrontShiny string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other struct{
			DreamWorld struct{
				FrontDefault string `json:"front_default"`
				FrontFemale string `json:"front_female"`
			} `json:"dream_world"`
			Home struct{
				FrontDefault string `json:"front_default"`
				FrontFemale string `json:"front_female"`
				FrontShiny string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct{
				FrontDefault string `json:"front_default"`
				FrontShiny string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct{
				BackDefault string `json:"back_default"`
				BackFemale string `json:"back_female"`
				BackShiny string `json:"back_shiny"`
				BackShinyFemale string `json:"back_shiny_female"`
				FrontDefault string `json:"front_default"`
				FrontFemale string `json:"front_female"`
				FrontShiny string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
	} `json:"sprites"` 
	Stats []struct{
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"stat"`
	} `json:"stats"` 
	Types []struct{
		Slot int `json:"slot"`
		Type struct{
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"type"`
	} `json:"types"` 
	Weight int `json:"weight"`
}

func (c *Client) getRequest(url string) (data []byte, status int, err error) {
	// Check cache
	if cacheValue, ok := c.cache.Get(url); ok {
		data = cacheValue
		status = 200
	} else {
		// Get request
		res, err := http.Get(url)
		status = res.StatusCode
		if err != nil {
			return []byte{}, status, err	
		}
		// Read from body stream
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return []byte{}, status, fmt.Errorf("Response failed with status code %v", res.StatusCode)
		}
		if err != nil {
			return []byte{}, status, err
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

	data, _, err := c.getRequest(url)
	if err != nil {
		return []string{}, "", "", err
	}
	
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
	data, status, err := c.getRequest(baseURL + "/v2/location-area/" + name)
	if err != nil {
		if status == 404 {
			err = fmt.Errorf(name + " is an invalid location.")
		}
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

func (c *Client) GetPokemon(name string) (pokemon Pokemon, err error) {
	data, status, err := c.getRequest(baseURL + "/v2/pokemon/" + name)
	if err != nil {
		if status == 404 {
			err = fmt.Errorf(name + " is an invalid pokemon.")
		}
		return Pokemon{}, err
	}

	jsonGetPokemon := Pokemon{}
	err = json.Unmarshal(data, &jsonGetPokemon)
	if err != nil {
		return 
	}

	return jsonGetPokemon, nil
}

type Client struct {
	cache pokecache.PokeCache
}

func NewClient(cacheTTLSeconds int) Client {
	return Client{cache: pokecache.New(cacheTTLSeconds)}
}