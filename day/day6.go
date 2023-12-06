package day

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meyerche/AdventOfCode2023/util"
)

type race struct {
	time int
	dist int64
}

func Day6Part1(data []string) int64 {
	fmt.Println("--- Begin Part 1 ---")

	races := getRaces(data)

	result := int64(1)
	for _, race := range races {
		result *= getWinScenarios(race)
	}
	return result
}

func Day6Part2(data []string) int64 {
	fmt.Println("--- Begin Part 1 ---")

	race := getBigRace(data)

	return getWinScenarios(race)
}

func getRaces(lines []string) []race {

	// times
	timeLine := lines[0]
	splitIdx := strings.IndexRune(timeLine, ':') + 1
	lineItems := strings.Fields(timeLine[splitIdx:])

	races := make([]race, len(lineItems))
	for i, lineItem := range lineItems {
		races[i].time = util.AtoiWrapper(lineItem)
	}

	// distance
	distanceLine := lines[1]
	splitIdx = strings.IndexRune(distanceLine, ':') + 1
	lineItems = strings.Fields(distanceLine[splitIdx:])

	for i, lineItem := range lineItems {
		d := util.AtoiWrapper(lineItem)
		races[i].dist = int64(d)
	}

	return races
}

func getBigRace(lines []string) race {
	var bigRace race

	// times
	timeLine := lines[0]
	splitIdx := strings.IndexRune(timeLine, ':') + 1
	timeString := strings.ReplaceAll(timeLine[splitIdx:], " ", "")
	bigRace.time = util.AtoiWrapper(timeString)

	// distance
	distanceLine := lines[1]
	splitIdx = strings.IndexRune(distanceLine, ':') + 1
	distString := strings.ReplaceAll(distanceLine[splitIdx:], " ", "")
	distInt, _ := strconv.ParseInt(distString, 10, 64)
	bigRace.dist = distInt

	return bigRace
}

func getWinScenarios(r race) int64 {
	count := int64(0)
	for t := 0; t <= r.time; t++ {
		travel := (r.time - t) * t
		if int64(travel) > r.dist {
			count++
		}
	}
	fmt.Println(r)
	fmt.Println(count)
	return count
}
