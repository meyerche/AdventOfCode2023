package day

import (
	"fmt"
	"strings"
	
	
	"github.com/meyerche/AdventOfCode2023/util"
)

func Day9Part1(lines []string) int {
	fmt.Println("--- Begin Part 1 ---")

	sum := 0
	
	for _, line := range lines {
		lineInt := getInts(line)
		sum += extrapolateOasis(lineInt, false)
		// fmt.Println(sum)
	}
	
	return sum
}

func Day9Part2(lines []string) int {
	fmt.Println("--- Begin Part 2 ---")

	sum := 0
	
	for _, line := range lines {
		lineInt := getInts(line)
		sum += extrapolateOasis(lineInt, true)
		// fmt.Println(sum)
	}
	
	return sum
}

func extrapolateOasis(line []int, goBack bool) int {
	newLine := make([]int, len(line) - 1)
	
	if allZero(line) {
		return 0
	}
	
	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]
		newLine[i-1] = diff
	}

	newValue := extrapolateOasis(newLine, goBack)
	
	var retVal int
	if goBack {
		retVal = line[0] - newValue
	} else {
		retVal = line[len(line)-1] + newValue
	}
	return retVal
}

func getInts(line string) []int {
	values := strings.Fields(line)
	newLine := make([]int, len(values))
	for i, val := range values {
		newLine[i] = util.AtoiWrapper(val)
	}
	return newLine
}

func allZero(line []int) bool {
	for _, val := range line {
		if val != 0 {
			return false
		}
	}
	return true
}