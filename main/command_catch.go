package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, arg ...string) error {
	if len(arg) != 1 {
		return errors.New("No Pokemon given")
	}

	pokemonName := arg[0]
	pokeapiclient := cfg.pokeapiclient

	resp, err := pokeapiclient.CatchPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("Error Retreiving pokemon %s", err)
	}

	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		fmt.Printf("Failed to catch %s\n", pokemonName)
		return nil
	}
	cfg.caughtPokemon[pokemonName] = resp
	fmt.Printf("Caught %s\n", pokemonName)
	return nil
}
