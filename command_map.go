package main

import "fmt"

func commandMap(cfg *config) error {
	locations := cfg.mapApi.GetNext()
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	locations := cfg.mapApi.GetPrevious()
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}
