package main

import (
	"fmt"
	"flag"
	
	"github.com/meyerche/AdventOfCode2023/util"
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
	
	
	fmt.Println(data)
}