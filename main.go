package main

import (
	"bufio"
	"errors"
	"fetch"
	"fmt"
	"os"
	"pokecache"
	"time"
)

//15 seconds cache
const cacheDuration = 1 * time.Second

type cliCommand struct {
	name        string
	description string
	callback    func(conf *fetch.Config_params, cache pokecache.Cache) error
}

func incrementConf(conf *fetch.Config_params) error {
	conf.Offset = conf.Offset + 20
	return nil
}

func decrementConf(conf *fetch.Config_params) error {
	if conf.Offset == 0 {
		return errors.New("pagination error")
	} else if conf.Offset == -20 {
		return errors.New("pagination error")
	} else {
		conf.Offset = conf.Offset - 20
	}
	return nil
}

func commandHelp(conf *fetch.Config_params, cache pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandLen(conf *fetch.Config_params, cache pokecache.Cache) error {
	fmt.Println(len(cache.Mapper))
	return nil
}

func commandExit(conf *fetch.Config_params, cache pokecache.Cache) error {
	os.Exit(0)
	return nil
}

func commandMap(conf *fetch.Config_params, cache pokecache.Cache) error {
	incrementConf(conf)
	resp, err := fetch.GET("https://pokeapi.co/api/v2/location-area/", conf, cache)
	if err != nil {
		return err
	}
	for value := range resp.Results {
		fmt.Println(resp.Results[value].Name)
	}
	return nil
}
func commandMapb(conf *fetch.Config_params, cache pokecache.Cache) error {
	err := decrementConf(conf)
	if err != nil {
		return err
	}
	resp, err := fetch.GET("https://pokeapi.co/api/v2/location-area/", conf, cache)
	if err != nil {
		return err
	}
	for value := range resp.Results {
		fmt.Println(resp.Results[value].Name)
	}
	return nil
}

func __init__() (fetch.Config_params, map[string]cliCommand) {
	conf := fetch.Config_params{
		Offset: -20,
		Limit:  20,
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
		"len": {
			name:        "len",
			description: "Displays the len of mapper",
			callback:    commandLen,
		},
	}
	return conf, commands
}

func main() {
	conf, commands := __init__()
	cache := pokecache.NewCache(cacheDuration)
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if value, ok := commands[scanner.Text()]; ok {
			err := value.callback(&conf, cache)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Error: Unknown command '" + string(scanner.Text()) + "'")
		}
		fmt.Print("Pokedex > ")
	}
}
