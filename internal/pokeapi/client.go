package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/appababba/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationAreaResponse, error) { // Added { here
	endpoint := fmt.Sprintf("/location-area/%s", locationAreaName)
	fullURL := baseURL + endpoint

	// check cache first
	if data, ok := c.cache.Get(fullURL); ok {
		locationAreaResp := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreaResp, nil
	}

	//if not in cache, make the request
	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Cache the response
	c.cache.Add(fullURL, data)

	// parse response
	locationAreaResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResp, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := fmt.Sprintf("/pokemon/%s", pokemonName)
	fullURL := baseURL + endpoint

	//check cache first
	if data, ok := c.cache.Get(fullURL); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// if not in cache, make the request
	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// cache the response
	c.cache.Add(fullURL, data)

	//parse response
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
