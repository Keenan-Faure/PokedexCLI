package pokefarm

import (
	"errors"
	"fetch"
	"fmt"
	"pokecache"
	"sync"
	"time"
)

const baseExp = 10
const expFirstForm = 150
const expSecondFrom = 400

type FarmPokemon struct {
	pokemon fetch.Pokemon
	baseExp int // 0=> initial | 1=>first | 2=>second
	time    time.Time
}

type PokeFarm struct {
	pokeFarm map[string]FarmPokemon
	Mux      *sync.Mutex
}

func CreatePokeFarm(interval time.Duration, cache pokecache.Cache) PokeFarm {
	pokeFarm := PokeFarm{
		pokeFarm: make(map[string]FarmPokemon),
		Mux:      &sync.Mutex{},
	}
	go pokeFarm.expLoop(interval, cache)
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

func (p *PokeFarm) expLoop(interval time.Duration, cache pokecache.Cache) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		p.addExpToFarm(interval, cache)
	}
}

func (p *PokeFarm) addExpToFarm(interval time.Duration, cache pokecache.Cache) {
	for _, value := range p.pokeFarm {
		url := "https://pokeapi.co/api/v2/pokemon-species/" + value.pokemon.Name
		if entry, ok := p.pokeFarm[value.pokemon.Name]; ok {
			entry.baseExp = p.calTotalExp(value.pokemon.Name)
			fmt.Println("I am here")
			fmt.Println(entry.pokemon.Name)
			if entry.baseExp > expFirstForm {
				poke_result, err := fetch.GETEvolID(url, value.pokemon.Name, &fetch.Config_params{}, cache)
				if err == nil {
					pokemon, err := p.GetPokemon(entry.pokemon.Name)
					if err == nil {
						p.Mux.Lock()
						defer p.Mux.Unlock()
						delete(p.pokeFarm, pokemon.pokemon.Name)
						p.AddPokemon(poke_result)
					}
				}
			} else if entry.baseExp > expSecondFrom {
				fetch.GETEvolID(url, value.pokemon.Name, &fetch.Config_params{}, cache)
			}
		}
	}
}

func (p *PokeFarm) CheckCurrExp() {
	if len(p.pokeFarm) > 0 {
		fmt.Println("Current Pokemon on PokeFarm: ")
		p.Mux.Lock()
		defer p.Mux.Unlock()
		for _, value := range p.pokeFarm {
			fmt.Printf("%s	|	%d", value.pokemon.Name, value.baseExp)
			fmt.Println("")
		}
	} else {
		fmt.Println("No pokemon at PokeFarm...")
	}
}
