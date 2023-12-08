package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/meyerche/AdventOfCode2023/day"
	"github.com/meyerche/AdventOfCode2023/util"
)

func main() {
	dayPtr := flag.Int("day", 0, "Advent day")
	filePtr := flag.String("file", "input1", "input or ex (example) plus.  default: input1 ")

	flag.Parse()

	if *dayPtr == 0 {
		fmt.Println("Must choose a day.")
		return
	}

	data := util.ReadFile(*dayPtr, *filePtr)
	runTheDay(*dayPtr, data)
}

func runTheDay(dayNum int, data []string) {
	switch dayNum {
	case 1:
		fmt.Println("Calibration 1 = " + strconv.Itoa(day.Day1Part1(data)))
		fmt.Println("Calibration 2 = " + strconv.Itoa(day.Day1Part2(data)))
	case 2:
		fmt.Println("Part 1 -- Sum of possible games = " + strconv.Itoa(day.Day2Part1(data)))
		fmt.Println("Part 2 -- Sum of set power = " + strconv.Itoa(day.Day2Part2(data)))
	case 3:
		fmt.Println("Part 1 -- " + strconv.Itoa(day.Day3Part1(data)))
	case 4:
		fmt.Println("*** Beginning Day 4 ***")
		fmt.Println("Total points = " + strconv.Itoa(day.Day4Part1(data)))
		fmt.Println("Total cards = " + strconv.Itoa(day.Day4Part2(data)))
	case 5:
		fmt.Println("*** Beginning Day 5 ***")
		fmt.Println("Nearest location = " + strconv.Itoa(day.Day5Part1(data)))
		fmt.Println("Nearest location = " + strconv.Itoa(day.Day5Part2(data)))
	case 6:
		fmt.Println("*** Beginning Day 6 ***")
		fmt.Println("Winning combos = " + strconv.FormatInt(day.Day6Part1(data), 10))
		fmt.Println("Big race combos = " + strconv.FormatInt(day.Day6Part2(data), 10))
	case 7:
		fmt.Println("*** Beginning Day 7 ***")
		fmt.Println("Total Winnings = " + strconv.Itoa(day.Day7Part1(data)))
		fmt.Println("Total Winnings = " + strconv.Itoa(day.Day7Part2(data)))
	case 8:
		fmt.Println("*** Beginning Day 8 ***")
		//fmt.Println("Steps = " + strconv.Itoa(day.Day8Part1(data)))
		fmt.Println("Ghost Steps = " + strconv.Itoa(day.Day8Part2(data)))
	}
}
