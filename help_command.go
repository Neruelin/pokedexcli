package main

import "fmt"

func helpCommand(commandState *CommandState) error {
	fmt.Println("Pokedex CLI is for looking up information about pokemon.")
	fmt.Println("Available Commands:")
	for name, command := range getCommands() {
		fmt.Println("\t", name, " - ", command.description)
	}
	return nil
}