package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config, args ...string) error {
	pokeapiclient := cfg.pokeapiclient

	resp, err := pokeapiclient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		log.Fatal("Shit broke")
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
