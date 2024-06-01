package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, arg ...string) error {
	if len(arg) != 1 {
		return errors.New("No LocationAreas given")
	}

	locationAreaName := arg[0]
	pokeapiclient := cfg.pokeapiclient

	resp, err := pokeapiclient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon in location area:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
