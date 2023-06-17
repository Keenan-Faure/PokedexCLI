package pokefarm

import (
	"errors"
	"fetch"
	"fmt"
	"pokecache"
	"sync"
	"time"
)

//create farm logger
//reports on the status of the farm

const baseExp = 5          //exp received every second
const expFirstForm = 1050  //exp required to evolve to 1st form
const expSecondFrom = 2500 //exp required to evolve to 2nd form

type FarmPokemon struct {
	pokemon fetch.Pokemon
	baseExp int
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
	return pokemon.pokemon, nil
}

func (p *PokeFarm) TransferPokemon(pokemonName string) (fetch.Pokemon, error) {
	_, err := p.GetPokemon(pokemonName)
	if err != nil {
		return fetch.Pokemon{}, err
	}
	pokemon, err := p.WithdrawPokemon(pokemonName)
	if err != nil {
		return fetch.Pokemon{}, err
	}
	fmt.Println("Withdrew", pokemonName, "from pokefarm")
	return pokemon, nil
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
			p.pokeFarm[value.pokemon.Name] = entry
			if entry.baseExp > expSecondFrom {
				poke_result, err := fetch.GETEvolID(url, value.pokemon.Name, &fetch.Config_params{}, cache)
				if err == nil {
					fmt.Println(entry.pokemon.Name + " is evolving...")
					pokemon, err := p.GetPokemon(entry.pokemon.Name)
					if err == nil {
						p.WithdrawPokemon(pokemon.pokemon.Name)
						fmt.Println(entry.pokemon.Name + " evolved into " + poke_result.Name)
						p.AddPokemon(poke_result)
					}
				}
			} else if entry.baseExp > expFirstForm {
				poke_result, err := fetch.GETEvolID(url, value.pokemon.Name, &fetch.Config_params{}, cache)
				if err == nil {
					fmt.Println(entry.pokemon.Name + " is evolving...")
					pokemon, err := p.GetPokemon(entry.pokemon.Name)
					if err == nil {
						p.WithdrawPokemon(pokemon.pokemon.Name)
						fmt.Println(entry.pokemon.Name + " evolved into " + poke_result.Name)
						p.AddPokemon(poke_result)
					}
				}
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
