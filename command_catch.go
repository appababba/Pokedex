package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/appababba/pokedexcli/internal/pokeapi"
)

// command handler for catch command
func HandleCatch(cfg *config, arg []string) error {
	if len(arg) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemonName := arg[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// call the catch function with the pokemon name
	err := Catch(cfg, pokemonName)
	if err != nil {
		return err
	}

	return nil

}

func calculateCatchThreshold(BaseExperience int) int {
	return max(1, 100-(BaseExperience/4))
}

var Pokedex = map[string]pokeapi.Pokemon{}

func Catch(cfg *config, pokemonName string) error {
	pokeData, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	//calculate catch threshold
	threshold := calculateCatchThreshold(pokeData.BaseExperience)

	//generate random chance
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := r.Intn(100) + 1

	//determine success
	if chance <= threshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		Pokedex[pokemonName] = pokeData
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}
