package pokeparty

import (
	"errors"
	"fetch"
	"sync"
	"time"
)

type PokeParty struct {
	pokeFarm map[string]fetch.Pokemon
	Mux      *sync.Mutex
}

func CreatePokeFarm(interval time.Duration) PokeParty {
	pokeFarm := PokeParty{
		pokeFarm: make(map[string]fetch.Pokemon),
		Mux:      &sync.Mutex{},
	}
	return pokeFarm
}

func (p *PokeParty) GetPokemon(pokemonName string) (fetch.Pokemon, error) {
	p.Mux.Lock()
	defer p.Mux.Unlock()
	pokemon, exist := p.pokeFarm[pokemonName]
	if exist {
		return pokemon, nil
	}
	return fetch.Pokemon{}, errors.New("no such pokemon at the pokefarm")
}

func (p *PokeParty) AddPokemon(pokemon fetch.Pokemon) {
	p.Mux.Lock()
	defer p.Mux.Unlock()
	p.pokeFarm[pokemon.Name] = pokemon
}

func (p *PokeParty) WithdrawPokemon(pokemonName string) (fetch.Pokemon, error) {
	pokemon, err := p.GetPokemon(pokemonName)
	if err != nil {
		return fetch.Pokemon{}, err
	}
	p.Mux.Lock()
	defer p.Mux.Unlock()
	delete(p.pokeFarm, pokemon.Name)
	return pokemon, nil
}
