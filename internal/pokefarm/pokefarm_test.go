package pokefarm

import (
	"fetch"
	"fmt"
	"testing"
	"time"
)

func TestNewFarmPokemon(t *testing.T) {
	pokemon := fetch.Pokemon{
		Name: "Example",
	}
	farmPokemon := newFarmPokemon(pokemon)
	if farmPokemon.pokemon.Name != pokemon.Name {
		t.Errorf("Expected 'Example' but found %s", farmPokemon.pokemon.Name)
	}
}

func TestCheckDuration(t *testing.T) {
	fmt.Println("Test Case 1 - Pokemon Exists")
	pokeFarm := CreatePokeFarm()
	pokemon := fetch.Pokemon{
		Name: "bulbasaur",
	}
	pokeFarm.AddPokemon(pokemon)
	time.Sleep(1 * time.Second)
	duration, _ := pokeFarm.checkDuration("bulbasaur")
	if duration < 0 {
		t.Errorf("Expected 'Duration > 0' but found 'Duration < 0'")
	}

	fmt.Println("Test Case 2 - Pokemon Does not Exist")
	pokeFarm = CreatePokeFarm()
	pokemon = fetch.Pokemon{
		Name: "catapie",
	}
	pokeFarm.AddPokemon(pokemon)
	time.Sleep(1 * time.Second)
	_, err := pokeFarm.checkDuration("bulbasaur")
	if err == nil {
		t.Errorf("Expected 'day care > we do not have that pokemon'")
	}
}
