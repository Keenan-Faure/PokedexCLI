package fetch

import (
	"fmt"
	"pokecache"
	"testing"
	"time"
)

func TestGET(t *testing.T) {
	const interval = 5 * time.Second
	cache := pokecache.NewCache(interval)
	querys := Config_params {
		Offset: 5,
		Limit: 1,
	}
	fmt.Println("Test Case 1 - Valid URL")
	_, ok := GET("https://pokeapi.co/api/v2/location-area/", &querys, cache)
	if(ok != nil) {
		t.Errorf(ok.Error())
	}

	fmt.Println("Test Case 2 - Invalid URL")
	_, ok = GET("", &querys, cache)
	fmt.Println(ok)
	if(ok.Error() != "undefined url") {
		t.Errorf("Expected 'undefined url' but found ''")
	}
}

func TestAddParams(t *testing.T) {
	fmt.Println("Test Case 1 - Actual Equals Expected")
	url := "http://example.com"
	params := Config_params{
		Offset: 5,
		Limit: 10,
	}
	expected := "http://example.com?offset=5&limit=10"
	actual := AddParams(url , &params)
	if(actual != expected) {
		t.Errorf("Expected " + expected + " but found " + actual)
	}

	fmt.Println("Test Case 1 - Actual Not Equal Expected")
	url = ""
	params = Config_params{
		Offset: 5,
		Limit: 10,
	}
	expected = "http://example.com?offset=5&limit=10"
	actual = AddParams(url , &params)
	if(actual == expected) {
		t.Errorf("Expected '' but found " + actual)
	}
}