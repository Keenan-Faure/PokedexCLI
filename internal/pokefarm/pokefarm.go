package pokefarm

import (
	"errors"
	"fetch"
	"sync"
	"time"
)

const baseExp = 5

type FarmPokemon struct {
	pokemon fetch.Pokemon
	time    time.Time
}

type PokeFarm struct {
	pokeFarm map[string]FarmPokemon
	Mux      *sync.Mutex
}

func CreatePokeFarm() PokeFarm {

	return PokeFarm{
		pokeFarm: make(map[string]FarmPokemon),
		Mux:      &sync.Mutex{},
	}
}

func newFarmPokemon(pokemon fetch.Pokemon) FarmPokemon {
	return FarmPokemon{
		pokemon: pokemon,
		time:    time.Now().UTC(),
	}
}

func (p *PokeFarm) checkDuration(pokemonName string) (time.Duration, error) {
	farmPokemon, err := p.GetPokemon(pokemonName)
	if err != nil {
		return time.Duration(0), err
	}
	time := time.Now().UTC().Sub(farmPokemon.time)
	return time, nil
}

func (p *PokeFarm) GetPokemon(pokemonName string) (FarmPokemon, error) {
	p.Mux.Lock()
	defer p.Mux.Unlock()
	pokemon, exist := p.pokeFarm[pokemonName]
	if exist {
		return pokemon, nil
	}
	return FarmPokemon{}, errors.New("day care > we do not have that pokemon")
}

func (p *PokeFarm) AddPokemon(pokemon fetch.Pokemon) {
	farmPokemon := newFarmPokemon(pokemon)
	p.Mux.Lock()
	defer p.Mux.Unlock()
	p.pokeFarm[pokemon.Name] = farmPokemon
}
