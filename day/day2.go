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

func Day2Part1(data []string) int {
	sum := 0

	for _, game := range data {
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

// func getGameNumber(game string) int {
// 	// index of colon
// 	idx1 := strings.Index(game, " ")
// 	idx2 := strings.Index(game, ":")
// 	value, _ := strconv.Atoi(game[idx1+1 : idx2])
// 	return value
// }
//
// func getSets(game string) []string {
// 	idx := strings.Index(game, ":")
// 	return strings.Split(game[idx+2:], "; ") //+2 accounts for the space after the colon
// }

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
