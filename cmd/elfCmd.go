package main

import (
	"fmt"
	"flag"
	"strconv"
	
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

func runTheDay(dayNum int, data []string) {
	switch dayNum {
	case 1:
		fmt.Println("Calibration 1 = " + strconv.Itoa(day.Day1Part1(data)))
		fmt.Println("Calibration 2 = " + strconv.Itoa(day.Day1Part2(data)))
	}
}