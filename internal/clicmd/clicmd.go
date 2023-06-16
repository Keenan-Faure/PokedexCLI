package clicmd

import (
	"errors"
	"fetch"
	"fmt"
	"math/rand"
	"os"
	"pokeball"
	"pokecache"
	"pokedex"
	"pokefarm"
	"pokeparty"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(
		name string,
		pokeball string,
		conf *fetch.Config_params,
		cache pokecache.Cache,
		pokedex *pokedex.Pokedex,
		pokeballs *pokeball.Pokeball,
		pokefarm *pokefarm.PokeFarm,
		pokeparty *pokeparty.PokeParty,
		pokeseen *fetch.SeenPoke) error
}

func commandChkPokeParty(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	pokeparty.CheckPartyPokemon()
	return nil
}

func commandTransferPokemon(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	pokemon, err := pokeparty.TransferPokemon(name)
	if err != nil {
		return err
	}
	fmt.Println("Transfering", name, "to pokefarm")
	pokefarm.AddPokemon(pokemon)
	fmt.Println(name, "was added to the pokefarm")
	return nil
}

func commandWithdrawPokeFarm(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	_, err := pokefarm.WithdrawPokemon(name)
	if err != nil {
		return nil
	}
	return nil
}

func commandCheckPokeFarm(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	pokefarm.CheckCurrExp()
	return nil
}

func commandAddPokeFarm(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	pokemon, exist := pokedex.GetPokemon(name)
	if !exist {
		return errors.New("you have not caught " + name)
	}
	fmt.Println(name, "was added to the pokefarm")
	pokefarm.AddPokemon(pokemon)
	return nil
}

func commandPokeballCount(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	fmt.Println("Available Pokeballs are: ")
	for key, value := range pokeballs.PokeBalls {
		fmt.Println(" - ", key, ": ", value)
	}
	return nil
}

func commandPokedex(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	if len(pokedex.Mapper) > 0 {
		for _, value := range pokedex.Mapper {
			fmt.Println(" - ", value.Name)
		}
	} else {
		fmt.Println("You have caught no pokemon")
	}
	return nil
}

func commandInspect(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	if name != "" {
		value, exist := pokedex.GetPokemon(name)
		if !exist {
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
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	if name != "" {
		_, err := pokeseen.GetPokemon(name)
		if err != nil {
			return err
		}
		pokemon, err := fetch.GETPokemon("https://pokeapi.co/api/v2/pokemon/"+name, conf, cache)
		if err != nil {
			return err
		}
		_, err = pokeballs.SubPokeball(pokeball)
		if err != nil {
			return err
		}
		_, exist := pokedex.GetPokemon(name)
		if exist {
			return errors.New("noo! I already have a " + pokemon.Name + " I dont need two...")
		}
		fmt.Println("Throwing a ", pokeball, "at ", name, "...")
		baseExperience := pokemon.BaseExperience
		rngNumber := rand.Intn(baseExperience * 2)
		catchChance := pokeballs.IncreaseChance(pokeball, baseExperience, rngNumber)
		if catchChance > baseExperience {
			fmt.Println(name, "was caught!")
			pokedex.AddPokemon(pokemon)
			fmt.Println("You may now inspect it with the inspect command.")
			result := pokeparty.AddPokemon(pokemon)
			if !result {
				pokefarm.AddPokemon(pokemon)
				return nil
			}
		} else {
			fmt.Println(name, "escaped!")
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
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	resp, err := fetch.GETExplore("https://pokeapi.co/api/v2/location-area/"+name, conf, cache)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + name + "...")
	fmt.Println("Found Pokemon:")
	for value := range resp.PokemonEncounters {
		fmt.Println(" - ", resp.PokemonEncounters[value].Pokemon.Name)
		pokeseen.AddPokemon(resp.PokemonEncounters[value].Pokemon.Name)
	}
	return nil
}

func commandHelp(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	fmt.Println("Welcome to Pokemon!")
	fmt.Println("Usage:")
	fmt.Println("Format:")
	fmt.Println("command <param1> <param2>              : Description")
	fmt.Println("")
	fmt.Println("help                                   : Displays a help message")
	fmt.Println("exit                                   : Exit the Pokedex")
	fmt.Println("map                                    : Displays the names of 20 location areas")
	fmt.Println("mapb                                   : Displays the previous 20 location areas")
	fmt.Println("explore      <area-name>               : Displays the Pokemon available in the specific area")
	fmt.Println("catch        <pokemon-name> <pokeball> : Attempts to Catch a pokemon using a pokeball")
	fmt.Println("inspect      <pokemon>                 : Inspect a caught pokemon's details")
	fmt.Println("pokedex                                : Displays caught pokemon")
	fmt.Println("poke-count                             : Displays the amount of pokemon caught")
	fmt.Println("poke-balls                             : Displays all held pokeballs")
	fmt.Println("farm-check                             : Displays the current pokemon growth rate")
	fmt.Println("farm-remove     <pokemon-name>         : Removes a Pokemon from the PokeFarm daycare")
	fmt.Println("party-check                            : Checks the pokemon in your party")
	fmt.Println("party-move   <pokemon-name>            : Removes a pokemon from your party and adds it to the Daycare")
	return nil
}

func commandPokedexCount(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	fmt.Println(len(pokedex.Mapper))
	return nil
}

func commandExit(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
	os.Exit(0)
	return nil
}

func commandMap(
	name string,
	pokeball string,
	conf *fetch.Config_params,
	cache pokecache.Cache,
	pokedex *pokedex.Pokedex,
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
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
	pokeballs *pokeball.Pokeball,
	pokefarm *pokefarm.PokeFarm,
	pokeparty *pokeparty.PokeParty,
	pokeseen *fetch.SeenPoke) error {
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

// helper functions
func incrementConf(conf *fetch.Config_params) error {
	conf.Offset = conf.Offset + 5
	return nil
}

func decrementConf(conf *fetch.Config_params) error {
	if conf.Offset == 0 {
		return errors.New("pagination error")
	} else if conf.Offset == -5 {
		return errors.New("pagination error")
	} else {
		conf.Offset = conf.Offset - 5
	}
	return nil
}

func Init() (pokeball.Pokeball, pokedex.Pokedex, fetch.Config_params, map[string]cliCommand) {
	pokeball := pokeball.InitPokeBalls()
	pokedex := pokedex.CreatePokedex()
	conf := fetch.Config_params{
		Offset: -5,
		Limit:  5,
	}
	commands := map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location areas",
			Callback:    commandMapb,
		},
		"pokecount": {
			Name:        "pokecount",
			Description: "Displays the amount of pokemon caught",
			Callback:    commandPokedexCount,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the Pokemon available in the specific area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to Catch a pokemon using a pokeball",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a caught pokemon's details",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays all caught pokemon",
			Callback:    commandPokedex,
		},
		"pokeballs": {
			Name:        "pokeballs",
			Description: "Displays all held pokeballs",
			Callback:    commandPokeballCount,
		},
		"add-farm": {
			Name:        "add-farm",
			Description: "Adds a Pokemon to the PokeFarm daycare",
			Callback:    commandAddPokeFarm,
		},
		"farm-check": {
			Name:        "farm-check",
			Description: "Displays the current pokemon growth rate",
			Callback:    commandCheckPokeFarm,
		},
		"rmv-farm": {
			Name:        "rmv-farm",
			Description: "Removes a Pokemon from the PokeFarm daycare",
			Callback:    commandWithdrawPokeFarm,
		},
		"party-check": {
			Name:        "party-check",
			Description: "Checks the pokemon in your party",
			Callback:    commandChkPokeParty,
		},
		"party-move": {
			Name:        "party-move",
			Description: "Removes a pokemon from your party and adds it to the Daycare",
			Callback:    commandTransferPokemon,
		},
	}
	return pokeball, pokedex, conf, commands
}
