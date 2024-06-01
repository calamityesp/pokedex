package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, arg ...string) error {
	if len(arg) != 1 {
		return errors.New("No Pokemon given")
	}

	pokemonName := arg[0]
	poke, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("Havvent caught the pokemon %s yet", pokemonName)
	}
	fmt.Println("Name:", poke.Name)
	fmt.Println("BaseExperience:", poke.BaseExperience)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weight:", poke.Weight)
	fmt.Println("Species:", poke.Species)
	return nil
}
