package main

import (
	"fmt"
	"github.com/calamityesp/pokedex/main/internal/pokeapi"
	"log"
)

func callbackMap(cfg *config) error {
	pokeapiclient := pokeapi.NewClient()

	resp, err := pokeapiclient.ListLocationAreas()
	if err != nil {
		log.Fatal("Shit broke")
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	return nil
}
