package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "       hello",
			expected: []string{"hello"},
		},
		{
			input:    "			hello",
			expected: []string{"hello"},
		},
		{
			input:    "hello           ",
			expected: []string{"hello"},
		},
		{
			input:    "hello			",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("TestCleanInput: Expected len(): %d but got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("TestCleanInput: input : '%s' \t expected: '%s' on index: %d \t got: '%s'", c.input, c.expected[i], i, actual[i])
			}
		}
	}
}

func TestUserInput(t *testing.T) {
	cases := []struct {
		input       string
		expectedStr string
		expectedOk  bool
	}{
		{
			input:       "Hello",
			expectedStr: "Hello",
			expectedOk:  true,
		},
		{
			input:       "",
			expectedStr: "",
			expectedOk:  false,
		},
	}

	for _, c := range cases {
		scanner := bufio.NewScanner(strings.NewReader(c.input))
		actual, actualOk := readUserInput(scanner)
		if len(c.expectedStr) != len(actual) {
			t.Errorf("TestUserInput: Expected len(): %d but got %d", len(c.expectedStr), len(actual))
		}
		if c.expectedOk != actualOk {
			t.Errorf("TestUserInput: Expected ok to be: %v but got %v", c.expectedOk, actualOk)
		}
		if c.expectedStr != actual {
			t.Errorf("TestUserInput: Expected: '%s' but got '%s'", c.expectedStr, actual)
		}
	}
}

func TestCaptureCommand(t *testing.T) {
	cases := []struct {
		input        string
		expected     string
		expectedRest []string
	}{
		{
			input:        "Hello World",
			expected:     "hello",
			expectedRest: []string{"World"},
		},
		{
			input:    "World",
			expected: "world",
		},
		{
			input:        "  hello  world  ",
			expected:     "hello",
			expectedRest: []string{"world"},
		},
		{
			input:        "CaTcH PiKaChu SECOND 123",
			expected:     "catch",
			expectedRest: []string{"PiKaChu", "SECOND", "123"},
		},
		{
			input:        "\t HeLLo\tWorld \n Test  ",
			expected:     "hello",
			expectedRest: []string{"World", "Test"},
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    " ",
			expected: "",
		},
		{
			input:    " \t\n ",
			expected: "",
		},
	}

	for _, c := range cases {
		actual, rest := captureCommand(c.input)

		if actual != c.expected {
			t.Errorf(
				"Input %q: expected command %q, got %q",
				c.input, c.expected, actual,
			)
		}

		if len(rest) != len(c.expectedRest) {
			t.Errorf(
				"Input %q: expected %d arguments %q, got %d arguments %q",
				c.input, len(c.expectedRest), c.expectedRest, len(rest), rest,
			)
			continue
		}

		for i, expected := range c.expectedRest {
			if rest[i] != expected {
				t.Errorf(
					"Input %q: expected argument at index %d to be %q, got %q",
					c.input, i, expected, rest[i],
				)
			}
		}
	}
}
