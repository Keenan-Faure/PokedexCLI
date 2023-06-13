package pokedex

import (
	"fetch"
	"testing"
)

func TestAddGetPokemon(t *testing.T) {
	pokedex := CreatePokedex()
	pokemon := fetch.Pokemon{
		Name: "example",
	}
	pokedex.AddPokemon(pokemon)
	_, ok := pokedex.GetPokemon("example")
	if(!ok) {
		t.Errorf("Expected true but found false")
	}
}