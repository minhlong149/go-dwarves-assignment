package main

import (
	"testing"
)

func TestCountRectangles(t *testing.T) {
	testCases := []struct {
		arrays   [][]int
		expected int
	}{
		{
			[][]int{},
			0,
		},
		{
			[][]int{
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			6,
		},
		{
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
			},
			1,
		},
		{
			[][]int{
				{1, 0, 1},
				{1, 0, 1},
			},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := countRectangles(testCase.arrays)
		if actual != testCase.expected {
			t.Errorf("countRectangles(%v) = %d; expected %d", testCase.arrays, actual, testCase.expected)
		}
	}
}
