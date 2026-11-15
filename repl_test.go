package main

import "testing"

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
			t.Errorf("Expected len(): %d but got %d", len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("TestCleanInput: input : '%s' \t expected: '%s' on index: %d \t got: '%s'", c.input, c.expected[i], i, actual[i])
			}
		}
	}
}
