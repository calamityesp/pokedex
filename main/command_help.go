package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome top the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	// loop through all of the cliCommands, and print out the name and descriptions
	for _, cmd := range getCommands() {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
