package helper

import (
	"strconv"
)

func ParseIntegers(strings []string) (integers []int, errors bool) {
	for _, strs := range strings {
		ints, parseError := strconv.Atoi(strs)
    if parseError != nil {
        errors = true
				return
    }
		integers = append(integers, ints)
	}
	return
}

func ParseFloats(strings []string) (floats []float64, errors bool) {
	for _, strs := range strings {
		flts, parseError := strconv.ParseFloat(strs, 64)
		if parseError != nil {
				errors = true
				return
		}
		floats = append(floats, flts)
	}
	return
}
