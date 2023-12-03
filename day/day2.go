package day

import (
	"strconv"
	"strings"
)

var LIMITS = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// type cubeSet struct {
// 	maxRed int
// 	maxGreen int
// 	maxBlue int
// }

func Day2Part1(games []string) int {
	sum := 0

	for _, game := range games {
		gameNumber := getGameNumber(game)
		validGame := true

		draws := getDraws(game)

		for _, draw := range draws {
			if !isValidDraw(draw) {
				// not valid draw
				validGame = false
				break
			}
		}

		if validGame {
			sum += gameNumber
		}
	}

	return sum
}

func Day2Part2(games []string) int {
	sum := 0

	for _, game := range games {
		// initialize to 1 so the power is not affected if there are no draws of a given color
		cubeSet := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		draws := getDraws(game)

		for _, draw := range draws {
			num, color := drawValues(draw)
			if num > cubeSet[color] {
				cubeSet[color] = num
			}
		}

		sum += cubeSetPower(cubeSet)
	}

	return sum
}

func getGameNumber(game string) int {
	// index of colon
	idx1 := strings.Index(game, " ")
	idx2 := strings.Index(game, ":")
	value, _ := strconv.Atoi(game[idx1+1 : idx2])
	return value
}

func getDraws(game string) []string {
	idx := strings.Index(game, ":")
	f := func(r rune) bool {
		return r == ';' || r == ','
	}

	return strings.FieldsFunc(game[idx+2:], f) //+2 accounts for the space after the colon
}

func isValidDraw(draw string) bool {
	draw = strings.Trim(draw, " ")
	parts := strings.Split(draw, " ")
	n, _ := strconv.Atoi(parts[0])

	return n <= LIMITS[parts[1]]
}

func drawValues(draw string) (int, string) {
	draw = strings.Trim(draw, " ")
	parts := strings.Split(draw, " ")
	n, _ := strconv.Atoi(parts[0])
	return n, parts[1]
}

func cubeSetPower(set map[string]int) int {
	return set["red"] * set["green"] * set["blue"]
}
