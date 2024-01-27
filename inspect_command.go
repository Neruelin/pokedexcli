package main

import (
	"fmt"
)

func validateInspect(args []string) bool {
	return len(args) == 1
}

func inspectCommand(commandState *CommandState, args []string) (err error) {
	if !validateInspect(args) {
		return fmt.Errorf("Invalid arguments to inspect command.")
	}
	
	if pokemon, ok := commandState.PokeList[args[0]]; ok {
		fmt.Printf("Name: %s\r\n", pokemon.Name)
		fmt.Printf("Height: %d\r\n", pokemon.Height)
		fmt.Printf("Weight: %d\r\n", pokemon.Weight)
		fmt.Printf("Stats:\r\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t- %s: %d\r\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\r\n")
		for _, poke_type := range pokemon.Types {
			fmt.Printf("\t- %s\r\n", poke_type.Type.Name)
		}
	} else {
		return fmt.Errorf(args[0] + " not found in pokedex.")
	}

	return 
}