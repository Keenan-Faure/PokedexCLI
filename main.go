package main

import (
	"bufio"
	"fmt"
	"internal/fetch"
	"os"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	fmt.Println("I am map")
	//https://pokeapi.co/api/v2/location-area?offset=0&limit=10
	return nil
}
func commandMapb() error {
	fmt.Println("I am mapb")
	return nil
}

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
	}

	fetch.Fetch("abc123", make(map[string]string))
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if value, ok := commands[scanner.Text()]; ok {
			value.callback()
		} else {
			fmt.Println("Error: Unknown command '" + string(scanner.Text()) + "'")
		}
		fmt.Print("Pokedex > ")
	}
}

//Try to fetch all the data from the API for that endpoint
	//try to make a struct that will accept the data (like a class)
	//Convert the response to a struct
	//transverse the struct in a nice format
		//use a function to return the key that you need
			//Function should transverse the struct and search for the key
//Try to filter the data that is being returned usin query param
	//Use a map to get the dat0a from the user
	//analyse the map
