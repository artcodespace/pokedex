package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello   world    ",
			expected: []string{"hello", "world"},
		},
		{
			input: "oneTwoThree",
			expected: []string{"onetwothree"},
		},
		{
			input: "oNe                         TWO three",
			expected: []string{"one", "two", "three"},
		},
	}

	for _,c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected output length %v; Got %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected output word %v; Got %v", expectedWord, word)
			}
		}
	}
}
