package main

import "fmt"

func helpCommand(commandState *CommandState, args []string) error {
	fmt.Printf("This Pokedex CLI tool is for listing information about locations in pokemon and what pokemon can be encountered there using https://pokeapi.co/.\r\n")
	fmt.Printf("Available Commands:\r\n")
	for _, command := range getCommands() {
		fmt.Printf("\t %s - %s\r\n", command.name, command.description)
	}
	return nil
}