package pokefarm

import (
	"fetch"
	"fmt"
	"pokecache"
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
	cache := pokecache.NewCache(15 * time.Second)
	pokeFarm := CreatePokeFarm(5*time.Second, cache)
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
	pokeFarm = CreatePokeFarm(5*time.Second, cache)
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

func TestGetPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - Pokemon Exists")
	cache := pokecache.NewCache(15 * time.Second)
	pokeFarm := CreatePokeFarm(5*time.Second, cache)
	pokemon := fetch.Pokemon{
		Name: "bulbasaur",
	}
	pokeFarm.AddPokemon(pokemon)
	_, ok := pokeFarm.GetPokemon("bulbasaur")
	if ok != nil {
		t.Errorf("Expected 'bulbasaur' but found ''")
	}

	fmt.Println("Test Case 2 - Pokemon does not exist")
	pokeFarm = CreatePokeFarm(5*time.Second, cache)
	pokemon = fetch.Pokemon{
		Name: "charmander",
	}
	pokeFarm.AddPokemon(pokemon)
	_, ok = pokeFarm.GetPokemon("bulbasaur")
	if ok == nil {
		t.Errorf("Expected '' but found 'bulbasaur'")
	}
}

func TestAddPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - One pokemon added to farm")
	cache := pokecache.NewCache(15 * time.Second)
	farmPokemon := CreatePokeFarm(5*time.Second, cache)
	pokemon := fetch.Pokemon{
		Name: "charizard",
	}
	farmPokemon.AddPokemon(pokemon)
	if len(farmPokemon.pokeFarm) != 1 {
		t.Errorf("Expected '1' but found %d", len(farmPokemon.pokeFarm))
	}
}

func TestWithdrawPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - Pokemon exists")
	cache := pokecache.NewCache(15 * time.Second)
	pokeFarm := CreatePokeFarm(5*time.Second, cache)
	pokemon := fetch.Pokemon{
		Name: "charmander",
	}
	pokeFarm.AddPokemon(pokemon)
	_, err := pokeFarm.WithdrawPokemon(pokemon.Name)
	if err != nil {
		t.Errorf("Expected error NOT to be nil")
	}
	if len(pokeFarm.pokeFarm) == 1 {
		t.Errorf("Expected '0' but found '1'")
	}

	fmt.Println("Test Case 2 - Pokemon does not exist")
	pokeFarm = CreatePokeFarm(5*time.Second, cache)
	pokemon = fetch.Pokemon{
		Name: "charmander",
	}
	pokeFarm.AddPokemon(pokemon)
	_, err = pokeFarm.WithdrawPokemon("bulbasaur")
	if err == nil {
		t.Errorf("Expected error to be nil")
	}
}

func TestCalTotalExp(t *testing.T) {
	cache := pokecache.NewCache(15 * time.Second)
	pokeFarm := CreatePokeFarm(3*time.Second, cache)
	pokemon := fetch.Pokemon{
		Name:           "charmander",
		BaseExperience: 65,
	}
	pokeFarm.AddPokemon(pokemon)
	time.Sleep(3*time.Second + time.Millisecond)
	exp := pokeFarm.calTotalExp("charmander")
	if exp != (3 * 10) {
		t.Errorf("Expected %d but found %d", (3 * 10), exp)
	}
	time.Sleep(6 * time.Second)
	pokeFarm.CheckCurrExp()
}
