package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Config_params struct {
	Offset string
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

func GET(url string, query_params *Config_params) (pokeloc, error){
	if(url != "") {
		url = add_query_params(url, query_params)
		resp, err := http.Get(url)
		if(err != nil) {
			return pokeloc{}, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if(err != nil) {
			return pokeloc{},err
		}
		result := pokeloc{}
		err_r := json.Unmarshal(body, &result)
		if(err_r != nil) {
			return pokeloc{}, err
		}
		return result, nil
	}
	return pokeloc{}, errors.New("undefined url")
}

func add_query_params(url string, query_params *Config_params) string {
	if(url != "" && query_params.Offset != "") {
		url = url + "?"
		url = url + "offset=" + query_params.Offset + "&"
		newLimit := fmt.Sprintf("%d", query_params.Limit)
		url = url + "limit=" + newLimit
		return url
	}
	return url
}
