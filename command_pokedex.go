package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
    // Check if any arguments were provided (should be none)
    if len(args) > 0 {
        return fmt.Errorf("usage: pokedex")
    }
    
    // Print the header
    fmt.Println("Your Pokedex:")
    
    // Check if Pokedex is empty
    if len(Pokedex) == 0 {
        fmt.Println("Empty! Go catch some Pokemon!")
        return nil
    }
    
    // Print each Pokemon name
    for name := range Pokedex {
        fmt.Printf(" - %s\n", name)
    }
    
    return nil
}
