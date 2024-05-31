package main

import (
	"fmt"
	"log"
)

func callbackMapp(cfg *config) error {
	client := cfg.pokeapiclient

	resp, err := client.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		log.Fatal("Shit really broke")
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
