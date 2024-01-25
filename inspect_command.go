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
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, poke_type := range pokemon.Types {
			fmt.Printf("\t- %s\n", poke_type.Type.Name)
		}
	} else {
		return fmt.Errorf(args[0] + " not found in pokedex.")
	}

	return 
}