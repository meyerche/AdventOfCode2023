package day

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type partNumber struct {
	value string
	index int
}
type partNumbers []partNumber

func Day3Part1(data []string) int {
	fmt.Println("*** Begin Day 3, Part 1 ***")

	sum := 0

	for i := 0; i < len(data); i++ {
		line := data[i]

		var prevLine string
		if i > 0 {
			prevLine = data[i-1]
		} else {
			prevLine = ""
		}

		var nextLine string
		if i < len(data)-1 {
			nextLine = data[i+1]
		} else {
			nextLine = ""
		}

		numbers := getNumbers(line)

		sum += processLine(numbers, line, prevLine, nextLine)
	}

	fmt.Println("*** End Day 3, Part 1 ***")

	return sum
}

func getNumbers(line string) partNumbers {
	var numbers partNumbers
	foundNumber := false
	var newNumber partNumber

	for i, value := range line {
		if !foundNumber && unicode.IsDigit(value) {
			foundNumber = true
			newNumber = partNumber{string(value), i}
		} else if foundNumber && unicode.IsDigit(value) {
			newNumber.value += string(value)
			if i == len(line)-1 {
				numbers = append(numbers, newNumber)
			}
		} else if foundNumber && !unicode.IsDigit(value) {
			foundNumber = false
			numbers = append(numbers, newNumber)
		}
	}

	return numbers
}

func processLine(nums partNumbers, line, prevLine, nextLine string) int {
	lineSum := 0

	for _, num := range nums {
		idx := num.index
		val := num.value
		isPart := false

		isPart = isPartNumber(idx, len(val), line)
		if len(prevLine) > 0 && !isPart {
			isPart = isPartNumber(idx, len(val), prevLine)
		}
		if len(nextLine) > 0 && !isPart {
			isPart = isPartNumber(idx, len(val), nextLine)
		}

		if isPart {
			lineSum += atoi(val)
		}
	}

	return lineSum
}

func isPartNumber(i, l int, line string) bool {
	f := func(c rune) bool {
		return !unicode.IsDigit(c) && c != '.'
	}

	var result bool
	if i == 0 {
		// number starts at the beginning of the line
		result = strings.ContainsFunc(line[:l+1], f)
	} else if (i + l) >= len(line)-1 {
		// number is at or one away from end of line so just check to the end of the line
		result = strings.ContainsFunc(line[i-1:], f)
	} else {
		result = strings.ContainsFunc(line[i-1:i+l+1], f)	
	}

	return result
}

// ASCII digits to integer.
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting part number to integer")
		i = 0
	}
	return i
}
