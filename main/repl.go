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

		cmd.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			name:        "Map",
			description: "Lists possible map location",
			callback:    callbackMap,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
