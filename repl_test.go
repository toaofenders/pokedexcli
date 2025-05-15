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
        // You might want to add a few more test cases
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        if len(actual) != len(c.expected) {
            t.Errorf("Expected output length %d but got %d for input %q", 
                     len(c.expected), len(actual), c.input)
            continue
        }

        for i := range actual {
            if i >= len(c.expected) {
                break
            }
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Expected word %q at position %d but got %q for input %q", 
                         expectedWord, i, word, c.input)
            }
        }
    }
}