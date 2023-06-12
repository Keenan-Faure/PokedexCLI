package main

import (
	"bufio"
	"fmt"
	"os"
	"pokecache"
	"strings"
	"time"
)

//15 seconds cache
const cacheDuration = 15 * time.Second

func main() {
	pokedex, conf, commands := __init__()
	cache := pokecache.NewCache(cacheDuration)
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command_params := strings.Fields(scanner.Text())
		command := command_params[0]
		explore_name := ""
		if(len(command_params) != 1) {
			explore_name = command_params[1]
		}
		if value, ok := commands[command]; ok {
			err := value.callback(explore_name, &conf, cache, &pokedex)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Error: Unknown command '" + string(scanner.Text()) + "'")
		}
		fmt.Print("Pokedex > ")
	}
}
