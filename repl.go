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
	callback    func(*config, ...string) error
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
		cmdStr, cmdArgs := captureCommand(text)
		cmd, ok := cfg.commands[cmdStr]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg, cmdArgs...)
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

func captureCommand(text string) (string, []string) {
	if text == "" {
		return "", nil
	}
	splitted := cleanInput(text)
	if len(splitted) == 0 {
		return "", nil
	}
	if len(splitted) > 1 {
		return strings.ToLower(splitted[0]), splitted[1:]
	}
	return strings.ToLower(splitted[0]), nil
}
