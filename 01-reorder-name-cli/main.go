package main

import (
	"fmt"
	"os"
)

var (
	firstName   string
	lastName    string
	middleName  string
	countryCode string
)

func reorderName(firstName, middleName, lastName, countryCode string,
) (reorderedName string, foundCountry bool) {
	type NameFormat int

	const (
		WesternOrder NameFormat = iota
		EasternOrder
	)

	countryMap := map[string]NameFormat{
		"US": WesternOrder,
		"VN": EasternOrder,
	}

	// Determine the name format based on the country code.
	countryNameFormat, foundCountry := countryMap[countryCode]
	if !foundCountry {
		return
	}

	// Reorder the name according to the determined format.
	switch countryNameFormat {

	case WesternOrder:
		reorderedName = firstName

		if middleName != "" {
			reorderedName += " " + middleName
		}

		reorderedName += " " + lastName

	case EasternOrder:
		reorderedName = lastName

		if middleName != "" {
			reorderedName += " " + middleName
		}

		reorderedName += " " + firstName
	}

	return
}

func main() {
	// Use the os package to access command line arguments
	args := os.Args[1:]

	// Parse command line arguments to extract the first name, last
	// name, middle name (if provided), and the country code.
	if len(args) < 3 {
		fmt.Printf("Error: Insufficient arguments\n")
		return
	}

	firstName = args[0]
	lastName = args[1]

	// Handle cases where the middle name is not provided.
	if len(args) == 3 {
		middleName = ""
		countryCode = args[2]
	} else {
		middleName = args[2]
		countryCode = args[3]
	}

	// Output the reordered name.
	if reorderedName, success := reorderName(firstName, middleName, lastName, countryCode); success {
		fmt.Printf("Output: %s\n", reorderedName)
	} else {
		fmt.Printf("Error: Country with code %s not found\n", countryCode)
	}
}
