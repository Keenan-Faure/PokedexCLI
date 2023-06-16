package pokeparty

import (
	"errors"
	"fetch"
	"fmt"
	"sync"
)

type PokeParty struct {
	pokeParty map[string]fetch.Pokemon
	Mux       *sync.Mutex
}

func CreatePokeParty() PokeParty {
	pokeFarm := PokeParty{
		pokeParty: make(map[string]fetch.Pokemon),
		Mux:       &sync.Mutex{},
	}
	return pokeFarm
}

func (p *PokeParty) GetPokemon(pokemonName string) (fetch.Pokemon, error) {
	p.Mux.Lock()
	defer p.Mux.Unlock()
	pokemon, exist := p.pokeParty[pokemonName]
	if exist {
		return pokemon, nil
	}
	return fetch.Pokemon{}, errors.New("no such pokemon in your party")
}

func (p *PokeParty) AddPokemon(pokemon fetch.Pokemon) bool {
	if len(p.pokeParty) >= 5 {
		fmt.Println("======")
		fmt.Println("Your party is currently full")
		fmt.Println("Tranferrring pokemon", pokemon.Name, "to pokefarm")
		return false
	}
	p.Mux.Lock()
	defer p.Mux.Unlock()
	p.pokeParty[pokemon.Name] = pokemon
	return true
}

func (p *PokeParty) TransferPokemon(pokemonName string) (fetch.Pokemon, error) {
	pokemon, err := p.GetPokemon(pokemonName)
	if err != nil {
		return fetch.Pokemon{}, err
	}
	p.Mux.Lock()
	defer p.Mux.Unlock()
	delete(p.pokeParty, pokemon.Name)
	return pokemon, nil
}

func (p *PokeParty) CheckPartyPokemon() {
	if len(p.pokeParty) == 0 {
		fmt.Println("You dont have any pokemon in your party")
	} else {
		fmt.Println("Current Pokemon in Party: ")
		p.Mux.Lock()
		defer p.Mux.Unlock()
		i := 0
		for _, value := range p.pokeParty {
			t1, t2 := getTypes(value)
			if t2 != "" {
				fmt.Printf("%d. ===> %s		===> %s, %s", i+1, value.Name, t1, t2)
			} else {
				fmt.Printf("%d. ===> %s		===> %s", i+1, value.Name, t1)
			}
			fmt.Println("")
			i++
		}
	}
}

// helper function
// Retrieves the types of the pokemon
func getTypes(pokemon fetch.Pokemon) (string, string) {
	if len(pokemon.Types) > 1 {
		return pokemon.Types[0].Type.Name, pokemon.Types[1].Type.Name
	}
	return pokemon.Types[0].Type.Name, ""
}
