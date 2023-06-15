package pokefarm

import (
	"errors"
	"fetch"
	"fmt"
	"sync"
	"time"
)

const baseExp = 3
const expFirstForm = 150
const expSecondFrom = 400

type FarmPokemon struct {
	pokemon fetch.Pokemon
	form    int
	baseExp int // 0=> initial | 1=>first | 2=>second
	time    time.Time
}

type PokeFarm struct {
	pokeFarm map[string]FarmPokemon
	Mux      *sync.Mutex
}

func CreatePokeFarm(interval time.Duration) PokeFarm {
	pokeFarm := PokeFarm{
		pokeFarm: make(map[string]FarmPokemon),
		Mux:      &sync.Mutex{},
	}
	go pokeFarm.expLoop(interval)
	return pokeFarm
}

func newFarmPokemon(pokemon fetch.Pokemon) FarmPokemon {
	return FarmPokemon{
		pokemon: pokemon,
		baseExp: 0,
		time:    time.Now().UTC(),
	}
}

func (p *PokeFarm) calTotalExp(pokemonName string) int {
	timeDuration, ok := p.checkDuration(pokemonName)
	if ok != nil {
		return 0
	}
	exp := int(timeDuration.Seconds()) * baseExp
	return exp
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
	return FarmPokemon{}, errors.New("no such pokemon at the pokefarm")
}

func (p *PokeFarm) AddPokemon(pokemon fetch.Pokemon) {
	farmPokemon := newFarmPokemon(pokemon)
	p.Mux.Lock()
	defer p.Mux.Unlock()
	p.pokeFarm[pokemon.Name] = farmPokemon
}

func (p *PokeFarm) WithdrawPokemon(pokemonName string) (fetch.Pokemon, error) {
	pokemon, err := p.GetPokemon(pokemonName)
	if err != nil {
		return fetch.Pokemon{}, err
	}
	p.Mux.Lock()
	defer p.Mux.Unlock()
	delete(p.pokeFarm, pokemon.pokemon.Name)
	fmt.Println("Successfully withdrew ", pokemonName)
	fmt.Println(pokemonName, " has ", pokemon.baseExp, " experience")
	return pokemon.pokemon, nil
}

func (p *PokeFarm) expLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		p.addExpToFarm(interval)
	}
}

func (p *PokeFarm) addExpToFarm(interval time.Duration) {
	for _, value := range p.pokeFarm {
		if entry, ok := p.pokeFarm[value.pokemon.Name]; ok {
			entry.baseExp = p.calTotalExp(value.pokemon.Name)
		}
	}
}

func (p *PokeFarm) CheckCurrExp() {
	fmt.Println("Current Pokemon on PokeFarm: ")
	p.Mux.Lock()
	defer p.Mux.Unlock()
	for _, value := range p.pokeFarm {
		fmt.Printf("%s	|	%d", value.pokemon.Name, value.baseExp)
		fmt.Println("")
	}
}
