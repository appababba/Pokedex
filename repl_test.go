package main

import "testing"

func TestCleanInput(t *testing.T) {
	// Define test cases
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	// Loop through test cases
	for _, c := range cases {
		actual := cleanInput(c.input) // Call the function being tested

		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("expected length %d, got %d", len(c.expected), len(actual))
		}

		// Check each word in the slice
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %s, got %s", expectedWord, word)
			}
		}
	}
}
