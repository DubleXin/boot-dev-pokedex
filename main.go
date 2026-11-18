package main

import "github.com/DubleXin/pokedexcli/api"

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
		},
		mapApi: api.MapApi{},
	}
	StartRepl(&cfg)

}
