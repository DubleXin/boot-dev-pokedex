package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config) error {

	if len(cfg.commands) == 0 {
		return errors.New("Commands map is empty")
	}
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range cfg.commands {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	return nil
}
