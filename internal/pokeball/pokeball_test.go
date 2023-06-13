package pokeball

import (
	"fmt"
	"testing"
)

func TestInitPokeBalls(t *testing.T) {
	pokeballs := InitPokeBalls()
	if(len(pokeballs.PokeBalls) != 4) {
		t.Errorf("Expected 4 but found %d", len(pokeballs.PokeBalls))
	}
}

func TestGetPokeball(t *testing.T) {
	fmt.Println("Test Case 1 - Valid Pokeball name")
	pokeballs := InitPokeBalls()
	_, ok := pokeballs.GetPokeball("poke-ball")
	if(ok != nil) {
		t.Errorf("Expected key 'poke-ball'")
	}

	fmt.Println("Test Case 2 - Invalid Pokeball name")
	_, ok = pokeballs.GetPokeball("Dive ball")
	if(ok == nil) {
		t.Errorf("Unexpected key 'Dive ball'")
	}
}

func TestSubPokeball(t *testing.T) {
	fmt.Println("Test Case 1 - Subtract valid Pokeball")
	pokeballs := InitPokeBalls()
	value, ok := pokeballs.SubPokeball("poke-ball")
	if(ok != nil) {
		t.Errorf("Expected 39 but found %d", value)
	}

	fmt.Println("Test Case 2 - Subtract Invalid Pokeball")
	_, ok = pokeballs.SubPokeball("Nest ball")
	if(ok == nil) {
		t.Errorf("Expected 'false' but found 'true'")
	}

	fmt.Println("Test Case 3 - Subtract higher pokeball count")
	pokeballs.PokeBalls["Master ball"] = 0
	_, ok = pokeballs.SubPokeball("master-ball")
	if(ok == nil) {
		t.Errorf("Expected 'false' but found 'true'")
	}
}

func TestIncreaseChance(t *testing.T) {
	fmt.Println("Test Case 1 - Using Pokeball")
	pokeballs := InitPokeBalls()
	baseExperience := 64
	rngNumber := 30
	newValue := pokeballs.IncreaseChance("poke-ball", baseExperience, rngNumber)
	if(newValue == rngNumber) {
		t.Errorf("Expected 31 but found %d", rngNumber)
	}

	fmt.Println("Test Case 2 - Using Great ball")
	newValue = pokeballs.IncreaseChance("great-ball", baseExperience, rngNumber)
	if(newValue == rngNumber) {
		t.Errorf("Expected 33 but found %d", rngNumber)
	}
	
	fmt.Println("Test Case 3 - Using Ultra ball")
	newValue = pokeballs.IncreaseChance("ultra-ball", baseExperience, rngNumber)
	if(newValue == rngNumber) {
		t.Errorf("Expected 36 but found %d", rngNumber)
	}

	fmt.Println("Test Case 4 - Using Master ball [100%]")
	newValue = pokeballs.IncreaseChance("master-ball", baseExperience, rngNumber)
	if(newValue == rngNumber) {
		t.Errorf("Expected 64 but found %d", rngNumber)
	}
}