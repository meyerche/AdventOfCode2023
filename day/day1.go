package day

import (
	"unicode"
)

func Day1Part1(data []string) int {
	calibration := 0
	for _, value := range data {
		calibration += (findFirstDigit(value) * 10) + findLastDigit(value)
	}
	
	return calibration
}

func findFirstDigit(word string) int {
	for _, ch := range word {
		if unicode.IsDigit(ch) {
			return int(ch)
		}
	}
	return 0
}

func findLastDigit(word string) int {
	runeWord := []rune(word)
	
	for i := len(runeWord)-1; i >= 0; i-- {
		if unicode.IsDigit(runeWord[i]) {
			return int(runeWord[i])
		}
	}
	return 0
}