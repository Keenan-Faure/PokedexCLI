package pokedex

import (
	"fetch"
	"sync"
)

type Pokedex struct {
	Mapper map[string]fetch.Pokemon
	Mux    *sync.Mutex
}

func CreatePokedex() Pokedex {
	return Pokedex{
		Mapper: make(map[string]fetch.Pokemon),
		Mux:    &sync.Mutex{},
	}
}

func (p *Pokedex) AddPokemon(pokemon fetch.Pokemon) {
	_, ok := p.GetPokemon(pokemon.Name)
	if !ok {
		p.Mux.Lock()
		defer p.Mux.Unlock()
		p.Mapper[pokemon.Name] = pokemon
		return
	}
}

func (p *Pokedex) GetPokemon(key string) (fetch.Pokemon, bool) {
	p.Mux.Lock()
	defer p.Mux.Unlock()
	_, ok := p.Mapper[key]
	if ok {
		return p.Mapper[key], true
	}
	return fetch.Pokemon{}, false
}
