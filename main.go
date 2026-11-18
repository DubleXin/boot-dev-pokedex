package main

import (
	"time"

	"github.com/DubleXin/pokedexcli/api"
	"github.com/DubleXin/pokedexcli/pokecache"
)

func main() {
	cfg := config{
		commands: map[string]cliCommand{
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"map": {
				name:        "map",
				description: "Next set of 20 locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "map back",
				description: "Previous set of 20 locations",
				callback:    commandMapb,
			},
			"explore": {
				name:        "explore",
				description: "Explore selected location",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Try to catch selected pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect known pokemon",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Display caught pokemons",
				callback:    commandPokedex,
			},
		},
		api: &api.PokeApi{
			Cache: pokecache.NewCache(time.Minute * 5),
		},
	}
	StartRepl(&cfg)

}
