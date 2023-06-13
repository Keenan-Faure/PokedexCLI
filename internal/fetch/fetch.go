package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pokecache"
)

type Config_params struct {
	Offset int
	Limit  int
}

func GET(url string, query_params *Config_params, cache pokecache.Cache) (pokeloc, error) {
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
