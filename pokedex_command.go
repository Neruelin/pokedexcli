package main

import "fmt"

func pokedexCommand(commandState *CommandState, args []string) (err error) {
	fmt.Printf("Your Pokedex:\r\n")
	for key, _ := range commandState.PokeList {
		fmt.Printf("\t- %s\r\n", key)
	}
	return 
}