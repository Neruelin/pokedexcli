package main

import (
	"fmt"
)

func mapCommand(commandState *CommandState, args []string) error {
	url := ""

	if commandState.Next != "" {
		url = commandState.Next
	}

	locations, previous, next, err := commandState.Client.GetLocations(url)
	if err != nil {
		return err
	}

	commandState.Previous = previous
	commandState.Next = next

	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}