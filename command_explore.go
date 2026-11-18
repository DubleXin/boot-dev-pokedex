package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please enter a location for exploration.")
		return nil
	}

	fmt.Printf("Exploring %s...\n", args[0])

	pokemons, err := cfg.api.GetExplorePokemons(args[0])
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if len(pokemons) == 0 {
		fmt.Println("No Pokemon found in this location.")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
