package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args ...string) error {
	names := cfg.api.GetCaughtPokemonNames()

	if len(names) == 0 {
		fmt.Println("You haven't caught any Pokemon yet. Try the catch command.")
		return nil
	}

	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
