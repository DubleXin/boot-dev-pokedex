package main

import "fmt"

func commandMap(cfg *config) error {
	locations := cfg.api.GetNext()
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	locations := cfg.api.GetPrevious()
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}
