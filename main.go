package main

import (
	"bufio"
	"clicmd"
	"fetch"
	"fmt"
	"os"
	"pokecache"
	"pokefarm"
	"pokeparty"
	"reflect"
	"strings"
	"time"
)

//DRY the fetch functions
//add a pokemon feel to the cli
//pokemon types like 🦅 for flying type etc

const cacheDuration = 15 * time.Second
const pokeFarmInterval = 20 * time.Second

func main() {
	pokeseen := fetch.CreateSeenPoke()
	fmt.Println(reflect.TypeOf(pokeseen))
	cache := pokecache.NewCache(cacheDuration)
	pokeparty := pokeparty.CreatePokeParty()
	pokefarm := pokefarm.CreatePokeFarm(pokeFarmInterval, cache)
	pokeballs, pokedex, conf, commands := clicmd.Init()
	fmt.Print("Pokemon > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command_params := strings.Fields(scanner.Text())
		command := command_params[0]
		explore_name := ""
		explore_pokeball := ""
		if len(command_params) == 2 {
			explore_name = command_params[1]
		} else if len(command_params) == 3 {
			explore_name = command_params[1]
			explore_pokeball = command_params[2]
		}
		if value, ok := commands[command]; ok {
			err := value.Callback(
				explore_name,
				explore_pokeball,
				&conf, cache,
				&pokedex,
				&pokeballs,
				&pokefarm,
				&pokeparty,
				&pokeseen,
			)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Error: Unknown command '" + string(scanner.Text()) + "'")
		}
		fmt.Print("Pokemon > ")
	}
}
