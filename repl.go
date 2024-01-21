package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/neruelin/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*CommandState) error
}

type CommandState struct {
	Next string
	Previous string
	Client pokeapi.Client
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: helpCommand,
		},
		"exit":{
			name: "exit",
			description: "Exits the CLI",
			callback: exitCommand,
		},
		"map": {
			name: "map",
			description: "Lists a page of 20 locations in the Pokemon world. Subsequent calls to map will return the next page of data",
			callback: mapCommand,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous page of 20 locations in the Pokemon world. Subsequent calls to mapb will return the next previous page of data",
			callback: mapbCommand,
		},
	}
}

func startRepl() {
	commandState := CommandState{}
	commandState.Client = pokeapi.NewClient()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		reader.Scan()
		
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&commandState)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}