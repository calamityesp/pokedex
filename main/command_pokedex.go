package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, arg ...string) error {
	fmt.Println("Printing the Pokedex")
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", poke.Name)
	}
	return nil
}
