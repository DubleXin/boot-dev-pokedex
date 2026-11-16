package main

import (
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
			expectedStr: "Failed to read input, try again.",
			expectedOk:  false,
		},
	}

	for _, c := range cases {
		actual, actualOk := readUserInput(strings.NewReader(c.input))
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
		input    string
		expected string
	}{
		{
			input:    "Hello World",
			expected: "hello",
		},
		{
			input:    "World",
			expected: "world",
		},
		{
			input:    "  hello  world  ",
			expected: "hello",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    " ",
			expected: "",
		},
	}
	for _, c := range cases {
		actual := captureCommand(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("TestCaptureCommand: Expected len(): %d but got %d", len(c.expected), len(actual))
		}
		if c.expected != actual {
			t.Errorf("TestCaptureCommand: Expected: '%s' but got '%s'", c.expected, actual)
		}
	}
}
