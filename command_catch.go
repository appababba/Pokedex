package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Pokemon struct {
	BaseExperience int `json:"base_experience"`
}

func Catch(pokemon string) {
	rand.Seed(time.Now().UnixNano()) // seed for randomness

	apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Failed to fetch pokemon data:", err)
		return
	}
	defer resp.Body.Close()

	// check if the response is valid
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch %s: HTTP Status %d\n", pokemon, resp.StatusCode)
		return
	}

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response body:", err)
		return
	}

	//parse the json extract base experiece
	var pokeData Pokemon
	err = json.Unmarshal(body, &pokeData)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	//Print the base experience to verify
	fmt.Printf("%s has a base experience of %d\n", pokemon, pokeData.BaseExperience)

	userNumber := rand.Intn(pokeData.BaseExperience) + 1
	randomCatch := rand.Intn(pokeData.BaseExperience) + 1

	if userNumber == randomCatch {
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped! Try again...\n", pokemon)
	}

}
