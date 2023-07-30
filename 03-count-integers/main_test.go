package main

import (
	"testing"
)

func TestCountInt (t *testing.T) {
	testCases := []struct {
		word string
		expected int
	}{
		{"a123bc34d8ef34", 3},
		{"420go69df", 2},
		{"0123456789", 1},
		{"abc", 0},
	}

	for _, testCase := range testCases {
		actual := numDifferentIntegers(testCase.word)
		if actual != testCase.expected {
			t.Errorf("numDifferentIntegers(%s) = %d; expected %d", testCase.word, actual, testCase.expected)
		}
	}
}
