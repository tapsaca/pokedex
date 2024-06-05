package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "input string",
			expected: []string {
				"input",
				"string",
			},
		},
		{
			input: "INPUT STRING",
			expected: []string {
				"input",
				"string",
			},
		},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("length '%v' is not equal to '%v'", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("input '%v' is not equal to '%v'", actualWord, expectedWord)
			}
		}
	}
}