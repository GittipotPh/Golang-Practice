package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// sort inputarrays with return the result as string / int /string
func customSort(item string) (string, int, string) {
	numPart := ""
	letterPart := ""

	for _, char := range item {
		if unicode.IsDigit(char) {
			numPart += string(char)
		} else if unicode.IsLetter(char) {
			letterPart += string(char)
		}
	}

	letterPart = strings.ToLower(letterPart)

	num, err := strconv.Atoi(numPart)
	if err != nil {
		// Handle the error
		return "", 0, item
	}

	return letterPart, num, item
}

func main() {
	inputArray := []string{"C1", "A9", "B8", "A1", "a5", "A5"}

	sort.Slice(inputArray, func(i, j int) bool {
		// Use all values returned by customSort
		letterI, numI, _ := customSort(inputArray[i])
		letterJ, numJ, _ := customSort(inputArray[j])

		//letter set to priority

		if letterI != letterJ {
			return letterI < letterJ
		} else if numI != numJ {
			return numI < numJ
		}
		return inputArray[i] < inputArray[j]
	})

	fmt.Println(inputArray)
}
