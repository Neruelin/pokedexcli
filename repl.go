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
	callback func(*CommandState, []string) error
}

type CommandState struct {
	Next string
	Previous string
	Client pokeapi.Client
	PokeList map[string]pokeapi.Pokemon
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
			description: "Lists a page of 20 location areas in the Pokemon world. Subsequent calls to map will return the next page of data",
			callback: mapCommand,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous page of 20 location areas in the Pokemon world. Subsequent calls to mapb will return the previous page of data",
			callback: mapbCommand,
		},
		"explore": {
			name: "explore <location-area>",
			description: "Lists the pokemon encountered in the provided <location-area>",
			callback: exploreCommand,
		},
		"catch": {
			name: "catch <pokemon>",
			description: "Attempts to catch and store the provided <pokemon>",
			callback: catchCommand,
		},
		"inspect": {
			name: "inspect <pokemon>",
			description: "Displays data about the provided <pokemon> if they have been caught and added to the pokedex",
			callback: inspectCommand,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists the names of the pokemon that have been caught",
			callback: pokedexCommand,
		},
	}
}

func startRepl() {
	commandState := CommandState{PokeList: map[string]pokeapi.Pokemon{}}
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
			err := command.callback(&commandState, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command. Try 'help' for information on what commands are available.")
			continue
		}
	}
}