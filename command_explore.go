package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	//validate args
	if len(args) != 1 {
		return fmt.Errorf("explore command needs exactly one argument")
	}
	locationAreaName := args[0]

	// let user know what we're doing
	fmt.Printf("Exploring %s...\n", locationAreaName)

	//Get the location area data
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	//print refults
	fmt.Println("Found pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
