package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pokecache"
	"sync"
)

type Config_params struct {
	Offset int
	Limit  int
}

type SeenPoke struct {
	seenPoke map[string]string
	Mux      *sync.Mutex
}

func CreateSeenPoke() SeenPoke {
	seenPoke := SeenPoke{
		seenPoke: make(map[string]string),
		Mux:      &sync.Mutex{},
	}
	return seenPoke
}

func (sp *SeenPoke) GetPokemon(pokemonName string, conf *Config_params, cache pokecache.Cache) (string, error) {
	sp.Mux.Lock()
	defer sp.Mux.Unlock()
	_, err := GETPokemon("https://pokeapi.co/api/v2/pokemon/"+pokemonName, conf, cache)
	if err != nil {
		return "", err
	}
	pokemon, exist := sp.seenPoke[pokemonName]
	if exist {
		return pokemon, nil
	}
	return "", errors.New("never seen a '" + pokemonName + "' yet\nplease explore your world...")
}

func (sp *SeenPoke) AddPokemon(pokemonName string) {
	sp.Mux.Lock()
	defer sp.Mux.Unlock()
	sp.seenPoke[pokemonName] = pokemonName
}

func (sp *SeenPoke) CountSeenPokemon() int {
	return len(sp.seenPoke)
}

func GETPokeLoc(url string, query_params *Config_params, cache pokecache.Cache) (pokeloc, error) {
	if url != "" {
		url = AddParams(url, query_params)
		cachedValue, exists := cache.Get(url)
		if exists {
			result := pokeloc{}
			err_r := json.Unmarshal(cachedValue, &result)
			if err_r != nil {
				return pokeloc{}, err_r
			}
			return result, nil
		}
		resp, err := http.Get(url)
		if err != nil {
			return pokeloc{}, err
		}
		defer resp.Body.Close()
		body, err_ := io.ReadAll(resp.Body)
		if err_ != nil {
			return pokeloc{}, err
		}
		cache.Add(url, body)
		result := pokeloc{}
		err_r := json.Unmarshal(body, &result)
		if err_r != nil {
			return pokeloc{}, err_r
		}
		return result, nil
	}
	return pokeloc{}, errors.New("undefined url")
}

func GETExplore(url string, query_params *Config_params, cache pokecache.Cache) (pokeExplore, error) {
	if url != "" {
		cachedValue, exists := cache.Get(url)
		if exists {
			result := pokeExplore{}
			err_r := json.Unmarshal(cachedValue, &result)
			if err_r != nil {
				return pokeExplore{}, err_r
			}
			return result, nil
		}
		resp, err := http.Get(url)
		if err != nil {
			return pokeExplore{}, err
		}
		defer resp.Body.Close()
		body, err_ := io.ReadAll(resp.Body)
		if err_ != nil {
			return pokeExplore{}, err
		}
		cache.Add(url, body)
		result := pokeExplore{}
		err_r := json.Unmarshal(body, &result)
		if err_r != nil {
			return pokeExplore{}, err_r
		}
		return result, nil
	}
	return pokeExplore{}, errors.New("undefined url")
}

func GETPokemon(url string, query_params *Config_params, cache pokecache.Cache) (Pokemon, error) {
	if url != "" {
		cachedValue, exists := cache.Get(url)
		if exists {
			result := Pokemon{}
			err_r := json.Unmarshal(cachedValue, &result)
			if err_r != nil {
				return Pokemon{}, err_r
			}
			return result, nil
		}
		resp, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()
		body, err_ := io.ReadAll(resp.Body)
		if err_ != nil {
			return Pokemon{}, err
		}
		cache.Add(url, body)
		result := Pokemon{}
		err_r := json.Unmarshal(body, &result)
		if err_r != nil {
			return Pokemon{}, err_r
		}
		return result, nil
	}
	return Pokemon{}, errors.New("undefined url")
}

func GETEvolID(url string, pokemonName string, query_params *Config_params, cache pokecache.Cache) (Pokemon, error) {
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()
		body, err_ := io.ReadAll(resp.Body)
		if err_ != nil {
			return Pokemon{}, err
		}
		result := PokeSpecies{}
		err_r := json.Unmarshal(body, &result)
		if err_r != nil {
			return Pokemon{}, err_r
		}
		evol, error_r := GETNextEvolution(result.EvolutionChain.URL, pokemonName, query_params)
		if error_r != nil {
			return Pokemon{}, error_r
		}
		next_form, err := GETPokemon(evol, query_params, cache)
		if err != nil {
			return Pokemon{}, err
		}
		return next_form, nil
	}
	return Pokemon{}, errors.New("undefined url")
}

func GETNextEvolution(url string, pokemonName string, query_params *Config_params) (string, error) {
	if url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err_ := io.ReadAll(resp.Body)
		if err_ != nil {
			return "", err
		}
		result := PokeEvolve{}
		err_r := json.Unmarshal(body, &result)
		if err_r != nil {
			return "", err_r
		}
		returns := evolution(pokemonName, result)
		if returns == "https://pokeapi.co/api/v2/pokemon/"+pokemonName {
			return "", errors.New("Pokemon cannot evolve again")
		}
		return evolution(pokemonName, result), nil
	}
	return "", errors.New("undefined url")
}

// helper function
// Adds the params to the url if it exists
func AddParams(url string, query_params *Config_params) string {
	if url != "" {
		url = url + "?"
		newOffset := fmt.Sprintf("%d", query_params.Offset)
		url = url + "offset=" + newOffset + "&"
		newLimit := fmt.Sprintf("%d", query_params.Limit)
		url = url + "limit=" + newLimit
		return url
	}
	return url
}

// determines which evolution form
// the POKeMON will advance towards
// returns a url that should be fetched
func evolution(pokemonName string, evolveMap PokeEvolve) string {
	//is it a first form
	if pokemonName == evolveMap.Chain.Species.Name {
		//move on to second form
		if len(evolveMap.Chain.EvolvesTo) > 0 {
			return "https://pokeapi.co/api/v2/pokemon/" + evolveMap.Chain.EvolvesTo[0].Species.Name
		}
		return "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	} else if pokemonName == evolveMap.Chain.EvolvesTo[0].Species.Name {
		//second form
		if len(evolveMap.Chain.EvolvesTo[0].EvolvesTo) > 0 {
			return "https://pokeapi.co/api/v2/pokemon/" + evolveMap.Chain.EvolvesTo[0].EvolvesTo[0].Species.Name
		}
		return "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	} else {
		//single form pokemon
		return "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	}
}
