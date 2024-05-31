package main

import (
	"time"

	"github.com/calamityesp/pokedex/main/internal/pokeapi"
)

type config struct {
	pokeapiclient       pokeapi.Client
	nextLocationAreaUrl *string // these are pointers so they can be nill
	prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiclient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}
