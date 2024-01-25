package main

import (
	"fmt"
	"math/rand"
)

func validateCatch(args []string) bool {
	return len(args) == 1
}

func catchCommand(commandState *CommandState, args []string) (err error) {
	if !validateCatch(args) {
		return fmt.Errorf("Invalid arguments to catch command.")
	}
	pokemon, err := commandState.Client.GetPokemon(args[0])
	if err != nil {
		return
	}

	roll := rand.Intn(pokemon.BaseExperience)

	if roll > 50 {
		fmt.Printf("%s escaped!\n", args[0])
	} else {
		fmt.Printf("%s was caught!\n", args[0])
		commandState.PokeList[args[0]] = pokemon
	}

	return 
}