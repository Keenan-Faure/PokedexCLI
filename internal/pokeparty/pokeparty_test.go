package pokeparty

import (
	"fetch"
	"fmt"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - Pokemon Exists")
	partyPokemon := CreatePokeParty()
	pokemon := fetch.Pokemon{
		Name: "bulbasaur",
	}
	partyPokemon.AddPokemon(pokemon)
	_, ok := partyPokemon.GetPokemon("bulbasaur")
	if ok != nil {
		t.Errorf("Expected 'bulbasaur' but found ''")
	}

	fmt.Println("Test Case 2 - Pokemon does not exist")
	partyPokemon = CreatePokeParty()
	pokemon = fetch.Pokemon{
		Name: "charmander",
	}
	partyPokemon.AddPokemon(pokemon)
	_, ok = partyPokemon.GetPokemon("bulbasaur")
	if ok == nil {
		t.Errorf("Expected '' but found 'bulbasaur'")
	}
}

func TestAddPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - One pokemon added to farm")
	partyPokemon := CreatePokeParty()
	pokemon := fetch.Pokemon{
		Name: "charizard",
	}
	partyPokemon.AddPokemon(pokemon)
	if len(partyPokemon.pokeParty) != 1 {
		t.Errorf("Expected '1' but found %d", len(partyPokemon.pokeParty))
	}
}

func TestTransferPokemon(t *testing.T) {
	fmt.Println("Test Case 1 - Pokemon exists")
	partyPokemon := CreatePokeParty()
	pokemon := fetch.Pokemon{
		Name: "charmander",
	}
	partyPokemon.AddPokemon(pokemon)
	_, err := partyPokemon.TransferPokemon(pokemon.Name)
	if err != nil {
		t.Errorf("Expected error NOT to be nil")
	}
	if len(partyPokemon.pokeParty) == 1 {
		t.Errorf("Expected '0' but found '1'")
	}

	fmt.Println("Test Case 2 - Pokemon does not exist")
	partyPokemon = CreatePokeParty()
	pokemon = fetch.Pokemon{
		Name: "charmander",
	}
	partyPokemon.AddPokemon(pokemon)
	_, err = partyPokemon.TransferPokemon("bulbasaur")
	if err == nil {
		t.Errorf("Expected error to be nil")
	}
}
