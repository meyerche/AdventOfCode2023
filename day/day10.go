package day

import (
	"fmt"
	"slices"
	"strings"
)

type coordinate struct {
	row int
	col int
}

var TILES = map[rune][]string{
	'|': {"north", "south"},
	'-': {"east", "west"},
	'L': {"north", "east"},
	'J': {"north", "west"},
	'7': {"south", "west"},
	'F': {"south", "east"},
}

var CONVERT_DIRECTION = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

func Day10Part1(lines []string) int {
	fmt.Println("--- Begin Part 1 ---")

	stepCount := 0
	startLoc := findStartTile(lines)
	fmt.Print("Found S at coordinates:  ")
	fmt.Println(startLoc)

	// make first move
	stepA, stepB := makeFirstMove(lines, startLoc.row, startLoc.col)
	locA := setCoordinates(stepA, startLoc)
	locB := setCoordinates(stepB, startLoc)
	stepCount++

	for locA != locB {
		// Path A
		stepA = moveDirection(rune(lines[locA.row][locA.col]), CONVERT_DIRECTION[stepA])
		locA = setCoordinates(stepA, locA)

		// Path B
		stepB = moveDirection(rune(lines[locB.row][locB.col]), CONVERT_DIRECTION[stepB])
		locB = setCoordinates(stepB, locB)

		stepCount++
	}

	return stepCount
}

// ********************************************************************
// This is all a bit of a mess. A result of going down multiple paths
// trying to get the logic right in order to identify interior points
// ********************************************************************
func Day10Part2(lines []string) int {
	fmt.Println("--- Begin Part 2 ---")

	points := initializeEmptyGrid(lines)
	pointStack := []coordinate{}

	startLoc := findStartTile(lines)

	// make first move
	stepA, stepB := makeFirstMove(lines, startLoc.row, startLoc.col)
	if (stepA == "north" || stepA == "south") && (stepB == "east" || stepB == "west") {
		points[startLoc.row][startLoc.col] = 2
	} else if (stepA == "east" || stepA == "west") && (stepB == "north" || stepB == "south") {
		points[startLoc.row][startLoc.col] = 2
	} else if (stepA == "north" || stepA == "south") && (stepB == "north" || stepB == "south") {
		points[startLoc.row][startLoc.col] = 1
	} else {
		points[startLoc.row][startLoc.col] = 0
	}
	for k, v := range TILES {
		if slices.Contains(v, stepA) && slices.Contains(v, stepB) {
			lines[startLoc.row] = strings.Replace(lines[startLoc.row], "S", string(k), 1)
		}
	}

	locA := setCoordinates(stepA, startLoc)
	pointStack = append(pointStack, locA)

	for locA != startLoc {
		// Path A
		stepA = moveDirection(rune(lines[locA.row][locA.col]), CONVERT_DIRECTION[stepA])
		if rune(lines[locA.row][locA.col]) == '|' {
			points[locA.row][locA.col] = 1
		} else if rune(lines[locA.row][locA.col]) == '-' {
			points[locA.row][locA.col] = 0
		} else {
			points[locA.row][locA.col] = 2
		}

		locA = setCoordinates(stepA, locA)
	}

	interiorPoints := countInterior(points, lines)
	fmt.Println(interiorPoints)
	// printPoints(points)

	return interiorPoints
}

func printPoints(points [][]int) {
	for _, row := range points {
		fmt.Println(row)
	}
}

func findStartTile(lines []string) coordinate {
	for row, line := range lines {
		col := strings.Index(line, "S")
		if col >= 0 {
			return coordinate{row, col}
		}
	}

	return coordinate{0, 0} // should never get here
}

func makeFirstMove(lines []string, row, col int) (string, string) {
	northRune := rune(lines[row-1][col])
	southRune := rune(lines[row+1][col])
	eastRune := rune(lines[row][col+1])
	westRune := rune(lines[row][col-1])

	// find the two starting paths
	var paths []string
	if slices.Contains(TILES[northRune], "south") {
		paths = append(paths, "north")
	}
	if slices.Contains(TILES[southRune], "north") {
		paths = append(paths, "south")
	}
	if slices.Contains(TILES[eastRune], "west") {
		paths = append(paths, "east")
	}
	if slices.Contains(TILES[westRune], "east") {
		paths = append(paths, "west")
	}

	return paths[0], paths[1]
}

func moveDirection(symbol rune, cameFrom string) (moveTo string) {
	// 	| is a vertical pipe connecting north and south.
	// 	- is a horizontal pipe connecting east and west.
	// 	L is a 90-degree bend connecting north and east.
	// 	J is a 90-degree bend connecting north and west.
	// 	7 is a 90-degree bend connecting south and west.
	// 	F is a 90-degree bend connecting south and east.
	// 	. is ground; there is no pipe in this tile.
	// 	S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	twoOptions := TILES[symbol]

	// 	fmt.Println(twoOptions)
	if twoOptions[0] == cameFrom {
		moveTo = twoOptions[1]
	} else {
		moveTo = twoOptions[0]
	}

	return moveTo
}

func setCoordinates(dir string, oldCoordinate coordinate) (newCoordinate coordinate) {
	switch dir {
	case "north":
		newCoordinate.row = oldCoordinate.row - 1
		newCoordinate.col = oldCoordinate.col
	case "south":
		newCoordinate.row = oldCoordinate.row + 1
		newCoordinate.col = oldCoordinate.col
	case "east":
		newCoordinate.row = oldCoordinate.row
		newCoordinate.col = oldCoordinate.col + 1
	case "west":
		newCoordinate.row = oldCoordinate.row
		newCoordinate.col = oldCoordinate.col - 1
	}

	return newCoordinate
}

func initializeEmptyGrid(lines []string) [][]int {
	points := make([][]int, len(lines))
	for i, line := range lines {
		points[i] = make([]int, len(line))
		for j := range line {
			points[i][j] = -1
		}
	}
	return points
}

func countInterior(grid [][]int, lines []string) int {
	count := 0
	isInterior := false
	lastCorner := '.'

	for i, row := range grid {
		// set tally segments to the right and left of each point
		fmt.Println(row)
		for j := range row {
			if grid[i][j] < 0 {
				if isInterior {
					count++
					fmt.Printf("row=%d col=%d -- count=%d \n", i, j, count)
				}
			} else if grid[i][j] == 1 {
				isInterior = !isInterior
			} else if grid[i][j] == 2 {
				corner := rune(lines[i][j])
				//fmt.Printf("row=%d col=%d -- corner=%s lastCorner=%s \n", i, j, string(corner), string(lastCorner))

				if cornerIsBorder(corner, lastCorner) {
					isInterior = !isInterior
					lastCorner = '.'
				} else {
					if lastCorner != '.' {
						lastCorner = '.'
					} else {
						lastCorner = rune(lines[i][j])
					}
				}
				//fmt.Println(isInterior)
			}
		}
		fmt.Println(count)
		isInterior = false
		lastCorner = '.'
	}

	return count
}

func cornerIsBorder(corner, lastCorner rune) bool {
	isBorder := false

	switch lastCorner {
	case 'F':
		if corner == 'J' {
			isBorder = true
		}
	case 'J':
		if corner == 'F' {
			isBorder = true
		}
	case '7':
		if corner == 'L' {
			isBorder = true
		}
	case 'L':
		if corner == '7' {
			isBorder = true
		}
	}

	fmt.Printf("lastCorner = %s, corner = %s, isBorder = %t\n", string(lastCorner), string(corner), isBorder)
	return isBorder
}

func sumIfPositive(values []int) int {
	sum := 0
	for _, val := range values {
		if val > 0 {
			sum += val
		}
	}
	return sum
}

func sumSegments(values []int) int {
	segments := []int{}
	sum := 0
	for _, val := range values {
		if val == 1 {
			segments = append(segments, val)
		} else if val == 2 {
			idx := slices.Index(segments, 2)
			if idx >= 0 {
				segments[idx] = 1
			} else {
				segments = append(segments, 2)
			}
		}
	}

	for _, v := range segments {
		sum += v
	}
	return sum
}
