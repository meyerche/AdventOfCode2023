package day

import (
	"strings"
	"unicode"
)

type calibrationValue struct {
	firstValue string
	firstIndex int
	lastValue  string
	lastIndex  int
}

func Day1Part1(lines []string) int {
	calibration := 0
	for _, line := range lines {
		calibration += (findFirstDigit(line) * 10) + findLastDigit(line)
	}

	return calibration
}

func Day1Part2(lines []string) int {
	calibration := 0
	for _, line := range lines {
		convertedLine := convertNumbers(line)
		calibration += (findFirstDigit(convertedLine) * 10) + findLastDigit(convertedLine)
	}

	return calibration
}

func findFirstDigit(word string) int {
	for _, ch := range word {
		if unicode.IsDigit(ch) {
			return int(ch - '0')
		}
	}
	return 0
}

func findLastDigit(word string) int {
	runeWord := []rune(word)

	for i := len(runeWord) - 1; i >= 0; i-- {
		if unicode.IsDigit(runeWord[i]) {
			return int(runeWord[i] - '0')
		}
	}
	return 0
}

func convertNumbers(word string) string {
	numbers := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var idx int
	value := calibrationValue{"0", len(word) + 1, "0", -1}

	for key, val := range numbers {
		idx = strings.Index(word, key)

		if idx >= 0 && idx < value.firstIndex {
			value.firstValue = val
			value.firstIndex = idx
		}

		idx = strings.LastIndex(word, key)

		if idx >= 0 && idx > value.lastIndex {
			value.lastValue = val
			value.lastIndex = idx
		}
	}

	if value.lastIndex >= 0 {
		word = word[:value.lastIndex] + value.lastValue + word[value.lastIndex:]
	}
	if value.firstIndex <= len(word) {
		word = word[:value.firstIndex] + value.firstValue + word[value.firstIndex:]
	}
	return word
}
