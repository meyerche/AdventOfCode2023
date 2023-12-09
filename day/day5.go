package day

import (
	"cmp"
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

	var sourceList, destinationList []int

	for i, values := range data {
		if i == 0 {
			sourceList = getInitialSeedList(values)
			destinationList = append(destinationList, sourceList...)
			// copy(destinationList, sourceList)
			continue
		}

		if strings.ContainsRune(values, ':') {
			copy(sourceList, destinationList)
			continue
		}

		if len(values) > 0 {
			mapToDestination(sourceList, destinationList, values)
		}
	}

	return slices.Min(destinationList)
}

func Day5Part2(data []string) int {
	fmt.Println("-- Begin part 2 --")

	var sourceRangeList, destinationRangeList []seedRange

	for i, values := range data {
		if i == 0 {
			sourceRangeList = getInitialSeedRangeList(values)
			continue
		}

		if strings.ContainsRune(values, ':') {
			destinationRangeList = append(destinationRangeList, sourceRangeList...)
			sourceRangeList = nil
			sourceRangeList = append(sourceRangeList, destinationRangeList...)
			destinationRangeList = nil
			continue
		}

		if len(values) > 0 {
			sourceRangeList, destinationRangeList = mapRangeToDestination(sourceRangeList, destinationRangeList, values)
		}
	}

	destinationRangeList = append(destinationRangeList, sourceRangeList...)
	return getNearestSeedRange(destinationRangeList)
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

	return seedsInt
}

func mapToDestination(src, dst []int, mapping string) {
	mapList := strings.Fields(mapping)

	destinationRangeStart, _ := strconv.Atoi(mapList[0])
	sourceRangeStart, _ := strconv.Atoi(mapList[1])
	rangeLength, _ := strconv.Atoi(mapList[2])

	for i, value := range src {
		if value >= sourceRangeStart && value < sourceRangeStart+rangeLength {
			dst[i] = destinationRangeStart + (value - sourceRangeStart)
		}
	}
}

func mapRangeToDestination(src, dst []seedRange, mapping string) ([]seedRange, []seedRange) {
	mapList := strings.Fields(mapping)
	destinationRangeStart, _ := strconv.Atoi(mapList[0])
	sourceRangeStart, _ := strconv.Atoi(mapList[1])
	rangeLength, _ := strconv.Atoi(mapList[2])

	// start -> before range
	// 		end -> before range  --- do nothing
	// 		end -> within range  --- split, write portion in range to destination, leave remainder in source
	// 		end -> after range   --- split, write middle portion in range to destination, leave first part in source, add additional range to source

	// start -> within range
	// 		end -> within range  --- write all range to destination, delete from source
	// 		end -> after range   --- write first portion in range to destination, leave extra range in source

	// start -> after range  --- do nothing, leave in src
	initSrcLen := len(src)

	for i := initSrcLen - 1; i >= 0; i-- {
		firstSeed := src[i].start
		lastSeed := src[i].start + src[i].rng - 1
		if firstSeed < sourceRangeStart {
			if lastSeed >= sourceRangeStart && lastSeed < sourceRangeStart+rangeLength {
				dst = append(dst, seedRange{destinationRangeStart, lastSeed - sourceRangeStart + 1})
				src[i].rng = sourceRangeStart - firstSeed
			} else if lastSeed >= sourceRangeStart+rangeLength {
				dst = append(dst, seedRange{destinationRangeStart, rangeLength})
				src[i].rng = sourceRangeStart - firstSeed
				src = append(src, seedRange{sourceRangeStart + rangeLength, lastSeed - (sourceRangeStart + rangeLength) + 1})
			}
		} else if firstSeed >= sourceRangeStart && firstSeed < sourceRangeStart+rangeLength {
			if lastSeed < sourceRangeStart+rangeLength {
				dst = append(dst, seedRange{firstSeed - (sourceRangeStart - destinationRangeStart), src[i].rng})
				src = slices.Delete(src, i, i+1)
			} else if lastSeed >= sourceRangeStart+rangeLength {
				dst = append(dst, seedRange{firstSeed - (sourceRangeStart - destinationRangeStart), sourceRangeStart + rangeLength - firstSeed})
				src[i].start = sourceRangeStart + rangeLength
				src[i].rng = lastSeed - (sourceRangeStart + rangeLength) + 1
			}
		}
	}
	
	return src, dst
}

func getNearestSeedRange(src []seedRange) int {
	m := func(a, b seedRange) int {
		return cmp.Compare(a.start, b.start)
	}

	minSeedRange := slices.MinFunc(src, m)
	return minSeedRange.start
}

func atoiWrapper(s string) int {
	// I wouldn't normally ignore potential errors like this.
	n, _ := strconv.Atoi(s)
	return n
}
