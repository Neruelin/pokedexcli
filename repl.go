package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/neruelin/pokedexcli/internal/pokeapi"
	"golang.org/x/term"
)

type cliCommand struct {
	name string
	description string
	callback func(*CommandState, []string) error
}

type CommandState struct {
	Next string
	Previous string
	Client pokeapi.Client
	PokeList map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: helpCommand,
		},
		"exit":{
			name: "exit",
			description: "Exits the CLI",
			callback: exitCommand,
		},
		"map": {
			name: "map",
			description: "Lists a page of 20 location areas in the Pokemon world. Subsequent calls to map will return the next page of data",
			callback: mapCommand,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous page of 20 location areas in the Pokemon world. Subsequent calls to mapb will return the previous page of data",
			callback: mapbCommand,
		},
		"explore": {
			name: "explore <location-area>",
			description: "Lists the pokemon encountered in the provided <location-area>",
			callback: exploreCommand,
		},
		"catch": {
			name: "catch <pokemon>",
			description: "Attempts to catch and store the provided <pokemon>",
			callback: catchCommand,
		},
		"inspect": {
			name: "inspect <pokemon>",
			description: "Displays data about the provided <pokemon> if they have been caught and added to the pokedex",
			callback: inspectCommand,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists the names of the pokemon that have been caught",
			callback: pokedexCommand,
		},
	}
}

func handleRawInput(history *[]string, historyIndex *int) (string, bool, error) {
	var currentBuffer []byte
	cursorIndex := 0
	for {
		var b = make([]byte, 1)
		_, err := os.Stdin.Read(b)
		char := b[0]
		if err != nil {
			return "", false, fmt.Errorf("Error reading from terminal: %s\r\n", err)
		}

		if (char == 3) { // ctrl-c
			fmt.Printf("\r\n")
			return "", true, nil
		}
		if char == 13 { // enter key
			fmt.Printf("\r\n")
			cursorIndex = 0
			*historyIndex = len(*history)
			break
		}
		if char == 127 { // backspace
			if cursorIndex > 0 {
				dist := len(currentBuffer) - cursorIndex
				fmt.Printf("\b")
				for idx := 0; idx <= dist; idx++ {
					fmt.Printf(" ")
				}
				for idx := 0; idx <= dist; idx++ {
					fmt.Printf("\b")
				}
				cursorIndex--
				if cursorIndex == 0 {
					currentBuffer = currentBuffer[1:]
				} else {
					currentBuffer = append(currentBuffer[:cursorIndex], currentBuffer[cursorIndex + 1:]...)
				}
				for idx := cursorIndex; idx < len(currentBuffer); idx++ {
					fmt.Printf("%c", currentBuffer[idx])
				}
				for idx := cursorIndex; idx < len(currentBuffer); idx++ {
					fmt.Printf("\b")
				}
			}
			
			continue
		}
		if char == 27 { // arrow sequence
			_, err := os.Stdin.Read(b)
			if err != nil {
				return "", false, fmt.Errorf("Error reading from terminal: %s\r\n", err)
			}
			if b[0] != 91 { // check for non-arrow key char sequences
				fmt.Printf("%c%c", char,b[0]) // replay the char sequence that has been consumed
				continue
			}
			_, err = os.Stdin.Read(b)
			if err != nil {
				return "", false, fmt.Errorf("Error reading from terminal: %s\r\n", err)
			}

			if b[0] == 'A' { // up arrow
				if *historyIndex == 0 {
					continue
				}
				inputLen := len(currentBuffer)
				for {
					if cursorIndex == inputLen {
						break
					}

					fmt.Printf(" ")
					cursorIndex++
				}
				for idx := 0; idx < inputLen; idx++ {
					fmt.Printf("\b ")
					fmt.Printf("\b")
				}
				cursorIndex = 0
				currentBuffer = []byte{}
				*historyIndex--
				fmt.Printf("%s", (*history)[*historyIndex])
				currentBuffer = []byte((*history)[*historyIndex])
				cursorIndex = len(currentBuffer)
				continue
			}

			if b[0] == 'B' { // down arrow
				if (*historyIndex > len(*history) - 1) {
					continue
				}

				inputLen := len(currentBuffer)
				for {
					if cursorIndex == inputLen {
						break
					}

					fmt.Printf(" ")
					cursorIndex++
				}
				for idx := 0; idx < inputLen; idx++ {
					fmt.Printf("\b ")
					fmt.Printf("\b")
				}
				cursorIndex = 0
				currentBuffer = []byte{}

				*historyIndex++

				if (*historyIndex == len(*history)) {
					continue	
				}

				currentBuffer = []byte((*history)[*historyIndex])
				for _, c := range currentBuffer {
					fmt.Printf("%c", c)
				}
				
				cursorIndex = len(currentBuffer)
				continue
			}

			if b[0] == 'D' { // left arrow
				if cursorIndex == 0 {
					continue
				}
				cursorIndex = cursorIndex - 1
			}

			if b[0] == 'C' { // right arrow
				if cursorIndex == len(currentBuffer) {
					continue
				}
				cursorIndex = cursorIndex + 1
			}

			fmt.Printf("%c%c%c", char, 91, b[0]) // replay full char sequence consumed
			
			continue
		}

		fmt.Printf("%c", char)

		if cursorIndex < len(currentBuffer) { // insert to middle of buffer
			currentBuffer = append(currentBuffer, currentBuffer[len(currentBuffer) - 1])
			for idx := len(currentBuffer) - 2; idx >= cursorIndex; idx-- {
				currentBuffer[idx + 1] = currentBuffer[idx]
			}
			currentBuffer[cursorIndex] = char
			cursorIndex += 1
			for idx := cursorIndex; idx < len(currentBuffer); idx++ {
				fmt.Printf("%c", currentBuffer[idx])
			}
			for idx := len(currentBuffer); idx > cursorIndex; idx-- {
				fmt.Printf("\b")
			}
			continue
		} else { // insert to end of buffer
			currentBuffer = append(currentBuffer, char)
			cursorIndex = len(currentBuffer)
		}
	}
	return string(currentBuffer), false, nil
}

func startRepl(cacheTTLSeconds int) {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	commandState := CommandState{PokeList: map[string]pokeapi.Pokemon{}}
	commandState.Client = pokeapi.NewClient(cacheTTLSeconds)
	
	history := []string{}
	historyIndex := 0

	for {
		fmt.Print("pokedexcli > ")

		inputString, shouldExit, err := handleRawInput(&history, &historyIndex)
		if shouldExit {
			return 
		}

		if err != nil {
			fmt.Printf("%s\r\n", err)
		}
		
		words := cleanInput(inputString)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&commandState, words[1:])
			if err != nil {
				fmt.Print(err)
				fmt.Printf("\r\n")
			}
			history = append(history, inputString)
			historyIndex = len(history)
			continue
		} else {
			fmt.Printf("Unknown command: '%s'. Try 'help' for information on what commands are available.\r\n", words[0])
			continue
		}
	}
}