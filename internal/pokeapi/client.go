package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/appababba/pokedexcli/internal/pokecache"
)

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
