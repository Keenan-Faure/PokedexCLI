package main

import (
	"bufio"
	"errors"
	"fetch"
	"fmt"
	"os"
	"pokecache"
	"sync"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *fetch.Config_params) error
}

func incrementConf(conf *fetch.Config_params) error{
	newLimit := fmt.Sprintf("%d", conf.Limit)
	conf.Offset = newLimit
	conf.Limit = conf.Limit + 20
	// fmt.Println("I am incrementor offset " + conf.Offset)
	// fmt.Println(conf.Limit)
	return nil
}

func decrementConf(conf *fetch.Config_params) error{
	if(conf.Limit == 0) {
		return errors.New("pagination error")
	} else if(conf.Limit == 20) {
		return errors.New("pagination error")
	} else {
		newLimit := fmt.Sprintf("%d", (conf.Limit - 40))
		conf.Offset = newLimit
		conf.Limit = conf.Limit - 20
	}
	// fmt.Println("I am decrementor offset " + conf.Offset)
	// fmt.Println(conf.Limit)
	return nil
}

func commandHelp(conf *fetch.Config_params) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit(conf *fetch.Config_params) error {
	os.Exit(0)
	return nil
}

func commandMap(conf *fetch.Config_params) error {
	incrementConf(conf)
	resp, err := fetch.GET("https://pokeapi.co/api/v2/location-area/", conf)
	if(err != nil) {
		return err
	}
	for value := range resp.Results {
		fmt.Println(resp.Results[value].Name)
	}
	return nil
}
func commandMapb(conf *fetch.Config_params) error {
	err := decrementConf(conf)
	if(err != nil) {
		return err
	}
	resp, err := fetch.GET("https://pokeapi.co/api/v2/location-area/", conf)
	if(err != nil) {
		return err
	}
	for value := range resp.Results {
		fmt.Println(resp.Results[value].Name)
	}
	return nil
}

func __init__() (fetch.Config_params, pokecache.Cache, map[string]cliCommand){
	conf := fetch.Config_params{
		Offset: "0",
		Limit: 0,
	}
	cache := pokecache.Cache {
		Mapper: make(map[string]pokecache.CacheEntry),
		Mux: &sync.Mutex{},
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
	return conf, cache, commands
}

func main() {
	conf, cache, commands := __init__()
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if value, ok := commands[scanner.Text()]; ok {
			err := value.callback(&conf)
			if(err != nil) {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Error: Unknown command '" + string(scanner.Text()) + "'")
		}
		fmt.Print("Pokedex > ")
	}
}
