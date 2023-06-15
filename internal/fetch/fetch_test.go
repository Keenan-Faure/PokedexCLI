package fetch

import (
	"fmt"
	"pokecache"
	"testing"
	"time"
)

func TestGET(t *testing.T) {
	const interval = 5 * time.Second
	cache := pokecache.NewCache(interval)
	querys := Config_params{
		Offset: 5,
		Limit:  1,
	}
	fmt.Println("Test Case 1 - Valid URL")
	_, ok := GET("https://pokeapi.co/api/v2/location-area/", &querys, cache)
	if ok != nil {
		t.Errorf(ok.Error())
	}

	fmt.Println("Test Case 2 - Invalid URL")
	_, ok = GET("", &querys, cache)
	fmt.Println(ok)
	if ok.Error() != "undefined url" {
		t.Errorf("Expected 'undefined url' but found ''")
	}
}

func TestAddParams(t *testing.T) {
	fmt.Println("Test Case 1 - Actual Equals Expected")
	url := "http://example.com"
	params := Config_params{
		Offset: 5,
		Limit:  10,
	}
	expected := "http://example.com?offset=5&limit=10"
	actual := AddParams(url, &params)
	if actual != expected {
		t.Errorf("Expected " + expected + " but found " + actual)
	}

	fmt.Println("Test Case 1 - Actual Not Equal Expected")
	url = ""
	params = Config_params{
		Offset: 5,
		Limit:  10,
	}
	expected = "http://example.com?offset=5&limit=10"
	actual = AddParams(url, &params)
	if actual == expected {
		t.Errorf("Expected '' but found " + actual)
	}
}

func TestAddPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - Length of 1")
	seenPoke := CreateSeenPoke()
	pokemon := "bulbasaur"
	seenPoke.AddPokemon(pokemon)
	if len(seenPoke.seenPoke) != 1 {
		t.Errorf("Expected '1' but found %d", len(seenPoke.seenPoke))
	}

	fmt.Println("Test Case 2 - Adding the same Pokemon")
	seenPoke = CreateSeenPoke()
	pokemon = "bulbasaur"
	seenPoke.AddPokemon(pokemon)
	seenPoke.AddPokemon(pokemon)
	if len(seenPoke.seenPoke) != 1 {
		t.Errorf("Expected '1' but found %d", len(seenPoke.seenPoke))
	}

	fmt.Println("Test Case 3 - Adding 3 different Pokemon")
	seenPoke = CreateSeenPoke()
	pokemon = "bulbasaur"
	pokemon2 := "charmander"
	pokemon3 := "chimchar"
	seenPoke.AddPokemon(pokemon)
	seenPoke.AddPokemon(pokemon2)
	seenPoke.AddPokemon(pokemon3)
	if len(seenPoke.seenPoke) != 3 {
		t.Errorf("Expected '1' but found %d", len(seenPoke.seenPoke))
	}
}

func TestGetSeenPoke(t *testing.T) {
	fmt.Println("Test Case 1 - Pokemon exists")
	seenPoke := CreateSeenPoke()
	pokemon := "bulbasaur"
	seenPoke.AddPokemon(pokemon)
	_, exist := seenPoke.GetPokemon("bulbasaur")
	if exist != nil {
		t.Errorf("Expected nil but found error")
	}

	fmt.Println("Test Case 2 - Pokemon does not exist")
	seenPoke = CreateSeenPoke()
	_, exist = seenPoke.GetPokemon("bulbasaur")
	if exist == nil {
		t.Errorf("Expected error but found none")
	}
}

func TestCountSeen(t *testing.T) {
	fmt.Println("Test Case 1 - len 1")
	seenPoke := CreateSeenPoke()
	pokemon := "bulbasaur"
	seenPoke.AddPokemon(pokemon)
	if seenPoke.CountSeenPokemon() != 1 {
		t.Errorf("Expected '1' but found %d", seenPoke.CountSeenPokemon())
	}

	fmt.Println("Test Case 1 - len 0")
	seenPoke = CreateSeenPoke()
	if seenPoke.CountSeenPokemon() != 0 {
		t.Errorf("Expected '0' but found %d", seenPoke.CountSeenPokemon())
	}
}
