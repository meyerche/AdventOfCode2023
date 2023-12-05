package day

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type seedRange struct {
	start int
	rng   int
}

func Day5Part1(data []string) int {
	fmt.Println("-- Begin part 1 --")
	// nearest := -1

	// almanacList := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	// idx := 0

	// var seedLocations []int
	// var source, destination string
	var sourceList, destinationList []int

	for i, values := range data {
		if i == 0 {
			sourceList = getInitialSeedList(values)
			destinationList = append(destinationList, sourceList...)
			//copy(destinationList, sourceList)
			continue
		}

		if strings.ContainsRune(values, ':') {
			// source = almanacList[idx]
			// destination = almanacList[idx+1]

			copy(sourceList, destinationList)
			continue
		}

		if len(values) > 0 {
			mapToDestination(sourceList, destinationList, values)
		}
	}

	return slices.Min(destinationList)
}

func getInitialSeedList(seeds string) []int {
	idx := strings.IndexRune(seeds, ':')
	seedsStr := strings.Fields(seeds[idx+1:])

	seedsInt := make([]int, len(seedsStr))
	for i, seed := range seedsStr {
		seedsInt[i], _ = strconv.Atoi(seed)
	}

	return seedsInt
}

func getInitialSeedRangeList(seeds string) []seedRange {
	idx := strings.IndexRune(seeds, ':')
	seedsStr := strings.Fields(seeds[idx+1:])

	seedsInt := make([]seedRange, len(seedsStr)/2)
	for i := 0; i < len(seedsStr); i += 2 {
		j := i / 2
		seedsInt[j] = seedRange{atoiWrapper(seedsStr[i]), atoiWrapper(seedsStr[i+1])}

	}

}

func mapToDestination(src, dst []int, mapping string) {
	mapList := strings.Fields(mapping)

	// fmt.Println(src)
	// fmt.Println(dst)
	// fmt.Println(mapping)

	destinationRangeStart, _ := strconv.Atoi(mapList[0])
	sourceRangeStart, _ := strconv.Atoi(mapList[1])
	rangeLength, _ := strconv.Atoi(mapList[2])

	for i, value := range src {
		if value >= sourceRangeStart && value < sourceRangeStart+rangeLength {
			dst[i] = destinationRangeStart + (value - sourceRangeStart)
		}
	}
}

func atoiWrapper(s string) int {
	// I wouldn't normally ignore potential errors like this.
	n, _ := strconv.Atoi(s)
	return n
}
