package main

import (
	"fmt"
	"flag"
	
	"github.com/meyerche/AdventOfCode2023/util"
	"github.com/meyerche/AdventOfCode2023/day"
)

func main() {
	dayPtr := flag.Int("day", 0, "Advent day")
	partPtr := flag.Int("part", 0, "test data (0), part 1 or 2")
	
	flag.Parse()
	
	if *dayPtr == 0 {
		fmt.Println("Must choose a day.")
		return
	}
	
	data := util.ReadFile(*dayPtr, *partPtr)
	runTheDay(*dayPtr, data)
}

func runTheDay(day int, data []string) {
	switch day {
	case 1:
		fmt.Println("Calibration = " + day.Day1Part1(data))
	}
}