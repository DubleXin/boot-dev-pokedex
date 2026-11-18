package main

import "fmt"

func commandMap(cfg *config, args ...string) error {
	locations := cfg.api.GetNextLocation()
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	locations := cfg.api.GetPreviousLocation()
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}
