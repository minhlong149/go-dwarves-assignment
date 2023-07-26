package main

import (
	"fmt"
	"strconv"
)

func numDifferentIntegers(word string) int {
	intMap := make(map[int]bool)

	var temp string
	for _, char := range word {
		if char >= '0' && char <= '9' {
			temp += string(char)
			continue
		}

		if temp != "" {
			integer, _ := strconv.Atoi(temp)
			intMap[integer] = true
			temp = ""
		}

	}

	if temp != "" {
		integer, _ := strconv.Atoi(temp)
		intMap[integer] = true
	}

	return len(intMap)
}

func main() {
	strWord := "1234bc34d1e00001f034A"
	countInt := numDifferentIntegers(strWord)
	fmt.Println(countInt)
}
