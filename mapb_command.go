package main

import (
	"fmt"
)

func mapbCommand(commandState *CommandState) error {
	if commandState.Previous == "" {
		return fmt.Errorf("mapb command error: No previous page\n")
	}

	locations, previous, next, err := commandState.Client.GetLocations(commandState.Previous)

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