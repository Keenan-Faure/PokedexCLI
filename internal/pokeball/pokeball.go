package pokeball

import (
	"errors"
	"sync"
)

type Pokeball struct {
	PokeBalls map[string]int
	Mux       *sync.Mutex
}

func InitPokeBalls() Pokeball {
	pokeballs := Pokeball{
		PokeBalls: map[string]int{
			"poke-ball":   40,
			"great-ball":  20,
			"ultra-ball":  10,
			"master-ball": 1,
		},
		Mux: &sync.Mutex{},
	}
	return pokeballs
}

func (p *Pokeball) GetPokeball(pokeball string) (int, error) {
	p.Mux.Lock()
	defer p.Mux.Unlock()
	value, ok := p.PokeBalls[pokeball]
	if ok {
		return value, nil
	}
	return 0, errors.New("pokeball " + pokeball + " not found")
}

func (p *Pokeball) SubPokeball(pokeball string) (int, error) {
	value, ok := p.GetPokeball(pokeball)
	if ok == nil {
		if value > 1 {
			p.Mux.Lock()
			defer p.Mux.Unlock()
			p.PokeBalls[pokeball] = (value - 1)
			return p.PokeBalls[pokeball], nil
		}
		return value, errors.New("you do not have enough " + pokeball)
	} else {
		if pokeball == "" {
			return value, errors.New("no pokeball selected")
		} else {
			return value, errors.New("you do not have a pokeball '" + pokeball + "'")
		}
	}
}

func (p *Pokeball) IncreaseChance(pokeball string, baseExperience int, rngNumber int) int {
	if pokeball == "poke-ball" {
		return rngNumber + int(baseExperience/40)
	} else if pokeball == "great-ball" {
		return rngNumber + int(baseExperience/20)
	} else if pokeball == "ultra-ball" {
		return rngNumber + int(baseExperience/10)
	} else if pokeball == "master-ball" {
		return baseExperience
	} else {
		return 0
	}
}
