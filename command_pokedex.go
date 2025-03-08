package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {

	if len(args) > 0 {
		return fmt.Errorf("usage: pokedex")
	}

	fmt.Println("Your Pokedex:")

	if len(Pokedex) == 0 {
		fmt.Println("Empty! Go catch some Pokemon!")
		return nil
	}

	for name := range Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
