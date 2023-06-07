package main

import (
	"bufio"
	"fetch"
	"fmt"
	"os"
)

type config struct {
	offset string
	limit  int
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

func commandHelp(conf config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit(conf config) error {
	os.Exit(0)
	return nil
}

func commandMap(conf config) error {
	resp, err := fetch.GET("https://pokeapi.co/api/v2/location-area/", query_params)
	if(err != nil) {
		fmt.Println(err.Error())
	}
	for value := range resp.Results {
		fmt.Println(resp.Results[value].Name)
	}
	return nil
}
func commandMapb(conf config) error {
	fmt.Println("I am mapb")
	return nil
}

func main() {
	conf := config {
		offset: "0",
		limit: 20,
	}
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    ,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit(conf),
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap(conf),
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb(conf),
		},
	}
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if value, ok := commands[scanner.Text()]; ok {
			value.callback(&conf)
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
