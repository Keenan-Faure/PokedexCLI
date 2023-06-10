package fetch

import (
	"fmt"
	"testing"
)

func TestGET(t *testing.T) {

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