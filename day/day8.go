package day

import (
	"fmt"
	"strings"
	"unicode"
)

func Day8Part1(lines []string) int {
	fmt.Println("--- Begin Part 1 ---")

	directions := lines[0]

	nodes := loadNodes(lines[2:])

	steps := navigate(directions, nodes)

	return steps
}

func Day8Part2(lines []string) int {
	fmt.Println("--- Begin Part 2 ---")

	directions := lines[0]
	nodes := loadNodes(lines[2:])

	steps := navigateLikeGhosts(directions, nodes)

	return steps
}

func loadNodes(lines []string) map[string][]string {
	nodes := make(map[string][]string)

	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsDigit(c)
		})

		nodes[fields[0]] = []string{fields[1], fields[2]}
	}

	return nodes
}

func navigate(dir string, nodes map[string][]string) int {
	count := 0

	loc := "AAA"
	for loc != "ZZZ" {
		for _, d := range dir {
			if d == 'L' {
				loc = nodes[loc][0]
			} else {
				loc = nodes[loc][1]
			}
			count++

			if loc == "ZZZ" {
				break
			}
		}
	}

	return count
}

func navigateFromLocation(directions string, loc string, nodes map[string][]string) int {
	count := 0

	for loc[2:] != "Z" {
		for _, dir := range directions {
			if dir == 'L' {
				loc = nodes[loc][0]
			} else {
				loc = nodes[loc][1]
			}
			count++

			if loc[2:] == "Z" {
				break
			}
		}
	}

	return count
}

func navigateLikeGhosts(directions string, nodes map[string][]string) int {
	locations := getStartingLocations(nodes)
	fmt.Println(locations)

	steps := make([]int, len(locations))

	for i, loc := range locations {
		steps[i] = navigateFromLocation(directions, loc, nodes)
	}

	// find least common multiple
	totalSteps := findLcm(steps)

	return totalSteps

	// *********************************
	// brute force method takes too long
	// *********************************
	// BPA ==> 11309
	// QCA ==> 15517
	// NDA ==> 20777
	// FDA ==> 19199
	// BVA ==> 17621
	// AAA ==> 16043
	// locations := []string{"BPA", "QCA", "NDA", "FDA", "BVA", "AAA"}

	// 	count := 0
	// 	finished := false
	// 	for !finished {
	// 		for _, dir := range directions {
	// 			finished = true
	//
	// 			// more one step for each current node location
	// 			for i, loc := range locations {
	// 				if dir == 'L' {
	// 					locations[i] = nodes[loc][0]
	// 				} else {
	// 					locations[i] = nodes[loc][1]
	// 				}
	//
	// 				if locations[i][2:] != "Z" {
	// 					finished = false
	// 				}
	// 				// else {
	// 				// 	fmt.Printf("%d - location: %d, node: %s, %s\n", count, i, locations[i], nodes[loc])
	// 				// }
	// 			}
	//
	// 			count++
	// 			if count%1000000 == 0 {
	// 				fmt.Println(count)
	// 			}
	//
	// 			if finished {
	// 				break
	// 			}
	// 		}
	// 	}
	//
	// 	return count
}

func getStartingLocations(nodes map[string][]string) []string {
	start := []string{}

	for node := range nodes {
		// fmt.Println(node)
		if node[2:] == "A" {
			start = append(start, node)
		}
	}

	return start
}

func findLcm(steps []int) int {
	totalSteps := steps[0]
	for i := 1; i < len(steps); i++ {
		totalSteps = lcm(totalSteps, steps[i])
	}
	return totalSteps
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		rem := a % b
		a = b
		b = rem
	}
	return a
}
