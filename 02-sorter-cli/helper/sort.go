package helper

import (
	"sort"
)

func SortIntegers(arrays []int) []int {
	sort.Ints(arrays)
	return arrays
}

func SortFloats(arrays []float64) []float64 {
	sort.Float64s(arrays)
	return arrays
}

func SortStrings(arrays []string) []string {
	sort.Strings(arrays)
	return arrays
}
