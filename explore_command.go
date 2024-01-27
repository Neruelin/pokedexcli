package main

import "fmt"

func validateExplore(args []string) bool {
	return len(args) == 1
}

func exploreCommand(commandState *CommandState, args []string) (err error) {
	if !validateExplore(args) {
		return fmt.Errorf("Invalid arguments to explore command.")
	}
	pokemonList, err := commandState.Client.GetLocation(args[0])
	if err != nil {
		return
	} else {
		fmt.Printf("Pokemon in %s:\r\n", args[0])
		l := len(pokemonList)
		for idx, name := range pokemonList {
			fmt.Print(name)
			if (idx != l - 1) {
				fmt.Print(", ")
			}
		}
		fmt.Printf("\r\n")
	}
	return
}