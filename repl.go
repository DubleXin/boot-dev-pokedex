package main

import (
	"bufio"
	"fmt"
	api "github.com/DubleXin/pokedexcli/api"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	api      *api.PokeApi
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func StartRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		text, ok := readUserInput(scanner)
		if !ok {
			break
		}
		if text == "" {
			continue
		}

		cmd, ok := cfg.commands[captureCommand(text)]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg)
		if err != nil {
			fmt.Print("Error ocured: %w\n", err.Error())
		}
	}
}

func readUserInput(scanner *bufio.Scanner) (scanned string, ok bool) {
	var text strings.Builder
	if scanner.Scan() {
		text.WriteString(scanner.Text())
		return text.String(), true
	}
	if scanner.Err() != nil {
		return "Failed to read input, try again.", false
	}
	return "", false

}

func captureCommand(text string) string {
	if text == "" {
		return ""
	}
	splitted := cleanInput(text)
	if len(splitted) == 0 {
		return ""
	}
	return strings.ToLower(splitted[0])
}
