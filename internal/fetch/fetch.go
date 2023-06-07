package fetch

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type pokeloc struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GET(url string, query_params map[string]string) (pokeloc, error){
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

func add_query_params(url string, query_params map[string]string) string {
	iterator := 0
	if(url != "" && query_params != nil) {
		for key := range query_params {
			if(iterator == 0) {
				url = url + "?"
				iterator ++
			}
			url = url + key + "=" + query_params[key] + "&"
		}
		url = strings.TrimSuffix(url, "&")
		return url
	}
	return url
}
