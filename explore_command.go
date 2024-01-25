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
		fmt.Println("Pokemon in " + args[0] + ":")
		l := len(pokemonList)
		for idx, name := range pokemonList {
			fmt.Print(name)
			if (idx != l - 1) {
				fmt.Print(", ")
			}
		}
		fmt.Println("")
	}
	return
}