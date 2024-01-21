package main

import (
	"fmt"
	"os"
)

func exitCommand(commandState *CommandState) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}