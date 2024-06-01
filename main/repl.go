package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	// Forever loop to run repl with go
	for {
		fmt.Print("pokedex >  ") // print without ln doesnt add a new line:w
		reader.Scan()
		text := reader.Text()

		// clean the input text
		cleaned := cleanInput(text)
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		// check for empty input and continue instead
		if len(cleaned) == 0 {
			continue
		}

		// parse input command
		command := cleaned[0]

		// check for command value
		cmd, ok := getCommands()[command]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		cmd.callback(cfg, args...)

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit cmd to leave the program",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists possible map locations",
			callback:    callbackMap,
		},
		"mapp": {
			name:        "mapp",
			description: "Lists previous map locations",
			callback:    callbackMapp,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemon in a given area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get basic information about a pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all of the pokemon you have caught",
			callback:    callbackPokedex,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
