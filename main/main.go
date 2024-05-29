package main

import (
	"fmt"
	"github.com/calamityesp/pokedex/main/internal/pokeapi"
	"log"
)

func main() {
	//	startRepl()
	pokeapiclient := pokeapi.NewClient()

	resp, err := pokeapiclient.ListLocationAreas()
	if err != nil {
		log.Fatal("Shit broke")
	}
	fmt.Println(resp)
}
