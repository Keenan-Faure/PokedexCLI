package main

import (
	"bufio"
	"clicmd"
	"fmt"
	"os"
	"pokecache"
	"strings"
	"time"
)

// 15 seconds cache
const cacheDuration = 15 * time.Second

func main() {
	pokeballs, pokedex, conf, commands := clicmd.Init()
	cache := pokecache.NewCache(cacheDuration)
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
			err := value.Callback(explore_name, explore_pokeball, &conf, cache, &pokedex, &pokeballs)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Error: Unknown command '" + string(scanner.Text()) + "'")
		}
		fmt.Print("Pokemon > ")
	}
}
