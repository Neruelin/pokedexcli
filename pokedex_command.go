package main

import "fmt"

func pokedexCommand(commandState *CommandState, args []string) (err error) {
	fmt.Println("Your Pokedex:")
	for key, _ := range commandState.PokeList {
		fmt.Printf("\t- %s\n", key)
	}
	return 
}