package main

import (
	"errors"
	"fetch"
	"fmt"
	"math/rand"
	"os"
	"pokecache"
	"pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(
		name string, 
		pokeball string, 
		conf *fetch.Config_params, 
		cache pokecache.Cache, 
		pokedex *pokedex.Pokedex,
		pokeballs *pokeball.Pokeball) error
}

func commandPokedex(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	for _, value := range pokedex.Mapper {
		fmt.Println(" - ", value.Name)
	}
	return nil
}

func commandInspect(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	if(name != "") {
		value, exist := pokedex.GetPokemon(name)
		if(!exist) {
			fmt.Println("you have not caught that pokemon")
		} else {
			fmt.Println("Name: ", value.Name)
			fmt.Println("Height: ", value.Height)
			fmt.Println("Weight: ", value.Weight)
			fmt.Println("Stats:")
			for _, key := range value.Stats {
				fmt.Println(" - ", key.Stat.Name, ": ", key.BaseStat)
			}
			fmt.Println("Types:")
			for _, key := range value.Types {
				fmt.Println(" - ", key.Type.Name)
			}
		}
	}
	return nil
}

func commandCatch(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	if(name != "") {
		resp, err := fetch.GETPokemon("https://pokeapi.co/api/v2/pokemon/" + name, conf, cache)
		if err != nil {
			return err
		}
		
		fmt.Println("Throwing a Pokeball at " + name + "...")
		baseExperience := resp.BaseExperience
		rngLimit := baseExperience * 2
		rngNumber := rand.Intn(rngLimit)
		if(rngNumber > baseExperience) {
			fmt.Println(name + " was caught!")
			pokedex.AddPokemon(resp)
			fmt.Println("You may now inspect it with the inspect command.")
		} else {
			fmt.Println(name + " escaped!")
		}
		return nil
	}
	return errors.New("undefined pokemon name ''")
}

func commandExplore(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	resp, err := fetch.GETExplore("https://pokeapi.co/api/v2/location-area/" + name, conf, cache)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + name + "...")
	fmt.Println("Found Pokemon:")
	for value := range resp.PokemonEncounters {
		fmt.Println(" - " + resp.PokemonEncounters[value].Pokemon.Name)
	}
	return nil
}

func commandHelp(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandLen(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	fmt.Println(len(cache.Mapper))
	return nil
}

func commandPokedexCount(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	fmt.Println(len(pokedex.Mapper))
	return nil
}

func commandExit(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
	os.Exit(0)
	return nil
}

func commandMap(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
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
func commandMapb(
	name string, 
	pokeball string, 
	conf *fetch.Config_params, 
	cache pokecache.Cache, 
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball) error {
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

//helper functions
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

func __init__() (pokeball.Pokeball, pokedex.Pokedex, fetch.Config_params, map[string]cliCommand) {
	pokeball := pokeball.InitPokeBalls()
	pokedex := pokedex.CreatePokedex()
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
		"pokecount": {
			name:        "pokecount",
			description: "Displays the amount of pokemon caught",
			callback:    commandPokedexCount,
		},
		"explore": {
			name:        "explore",
			description: "Displays the Pokemon available in the specific area",
			callback:    commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempts to Catch a pokemon using a pokeball",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a caught pokemon's details",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays all caught pokemon",
			callback: commandPokedex,
		},
	}
	return pokeball, pokedex, conf, commands
}