package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Hello   World  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "Charizard bulbasaur PIKACHU",
			expected: []string{"charizard", "bulbasaur", "pikachu"},
		}, {
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word %d to be %s, got %s", i, expectedWord, word)
			}
		}
	}
}
