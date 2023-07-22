package main

import (
	"flag"
	"fmt"

	"github.com/minhlong149/go-dwarves-assignment/02-sorter-cli/helper"
)

var (
	intFlag    bool
	floatFlag  bool
	stringFlag bool
)

func init() {
	flag.BoolVar(&intFlag, "int", false, "Sort integers")
	flag.BoolVar(&floatFlag, "float", false, "Sort floats")
	flag.BoolVar(&stringFlag, "string", false, "Sort strings")
	flag.Parse()
}

func main() {
	switch {
	case intFlag:
		integers, errors := helper.ParseIntegers(flag.Args())
		if errors {
			fmt.Println("Error: Not all arguments are integers")
			return
		}
		fmt.Println("Output:", helper.SortIntegers(integers))

	case floatFlag:
		floats, errors := helper.ParseFloats(flag.Args())
		if errors {
			fmt.Println("Error: Not all arguments are floats")
			return
		}
		fmt.Println("Output:", helper.SortFloats(floats))

	case stringFlag:
		fmt.Println("Output:", helper.SortStrings(flag.Args()))

	default:
		fmt.Println("No flag specified. Please select a type to sort.")
	}
}
