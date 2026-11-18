package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 || strings.TrimSpace(args[0]) == "" {
		fmt.Println("Please enter a Pokemon name to catch.")
		return nil
	}

	name := strings.ToLower(strings.TrimSpace(args[0]))

	pokemon, err := cfg.api.GetPokemon(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if _, ok := cfg.api.GetCaughtPokemon(pokemon.Name); ok {
		fmt.Printf("%s has already been caught!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := math.Max(
		0.1,
		100.0/(100.0+float64(pokemon.BaseExperience)),
	)

	if rand.Float64() >= chance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	cfg.api.AddCaughtPokemon(pokemon)
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
