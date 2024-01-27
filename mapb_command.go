package main

import (
	"fmt"
)

func mapbCommand(commandState *CommandState, args []string) error {
	if commandState.Previous == "" {
		return fmt.Errorf("mapb command error: No previous page")
	}

	locations, previous, next, err := commandState.Client.GetLocations(commandState.Previous)

	if err != nil {
		return err
	}

	commandState.Previous = previous
	commandState.Next = next

	for _, location := range locations {
		fmt.Printf("%s\r\n", location)
	}

	return nil
}