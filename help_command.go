package main

import "fmt"

func helpCommand(commandState *CommandState, args []string) error {
	fmt.Println("This Pokedex CLI tool is for listing information about locations in pokemon and what pokemon can be encountered there using https://pokeapi.co/.")
	fmt.Println("Available Commands:")
	for _, command := range getCommands() {
		fmt.Printf("\t %s - %s\n", command.name, command.description)
	}
	return nil
}