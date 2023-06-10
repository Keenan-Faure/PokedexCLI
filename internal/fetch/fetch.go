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

type pokeloc struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GET(url string, query_params *Config_params, cache pokecache.Cache) (pokeloc, error){
	if(url != "") {
		url = AddParams(url, query_params)
		cachedValue, exists := cache.Get(url)
		if(exists) {
			result := pokeloc{}
			err_r := json.Unmarshal(cachedValue, &result)
			if(err_r != nil) {
				return pokeloc{}, err_r
			}
			return result, nil
		}
		resp, err := http.Get(url)
		if(err != nil) {
			return pokeloc{}, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if(err != nil) {
			return pokeloc{},err
		}
		cache.Add(url, body)
		result := pokeloc{}
		err_r := json.Unmarshal(body, &result)
		if(err_r != nil) {
			return pokeloc{}, err_r
		}
		return result, nil
	}
	return pokeloc{}, errors.New("undefined url")
}

func AddParams(url string, query_params *Config_params) string {
	if(url != "") {
		url = url + "?"
		newOffset := fmt.Sprintf("%d", query_params.Offset)
		url = url + "offset=" + newOffset + "&"
		newLimit := fmt.Sprintf("%d", query_params.Limit)
		url = url + "limit=" + newLimit
		return url
	}
	return url
}
