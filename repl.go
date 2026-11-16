package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func StartRepl() {
	for {
		fmt.Print("Pokedex > ")
		text, ok := readUserInput(os.Stdin)
		if !ok {
			text, ok = readUserInput(os.Stdin)
			continue
		}

		fmt.Print("Your command was: ", captureCommand(text), "\n")
	}
}

func readUserInput(rd io.Reader) (scanned string, ok bool) {

	scanner := bufio.NewScanner(rd)
	var text strings.Builder
	for scanner.Scan() {
		scanend := scanner.Text()
		if scanend != "" {
			text.WriteString(scanend)
			break
		}
	}
	if scanner.Err() != nil || text.String() == "" {
		return "Failed to read input, try again.", false
	}
	return text.String(), true

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
