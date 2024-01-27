package main

import (
	"fmt"
	"os"
)

func exitCommand(commandState *CommandState, args []string) error {
	fmt.Printf("Exiting...\r\n")
	os.Exit(0)
	return nil
}