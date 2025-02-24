package main

import "fmt"

func inspectPokemon(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	pokemonName := args[0]

	//check if the pokemon exists in the Pokedex map
	pokemon, exists := Pokedex[pokemonName]
	if !exists {
		fmt.Println("you have not caught this pokemon!")
		return nil
	}

	// if caught, print the pokemon's details
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	- %s: %d\n", stat.StatInfo.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("	-%s\n", pType.Type.Name)
	}

	return nil
}
