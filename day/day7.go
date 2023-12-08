package day

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/meyerche/AdventOfCode2023/util"
)

type handStruct struct {
	hand string
	bid  int
}
type handStructs []handStruct

var cardStrength = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func Day7Part1(lines []string) int {
	fmt.Println("--- Begin Part 1 ---")

	handType := groupInputIntoHandTypes(lines, false)

	return calculateWinnings(handType)

}

func Day7Part2(lines []string) int {
	fmt.Println("--- Begin Part 2 ---")

	// Joker is lowest value card
	cardStrength['J'] = 0

	// Replace joker
	handType := groupInputIntoHandTypes(lines, true)

	return calculateWinnings(handType)
}

func groupInputIntoHandTypes(lines []string, joker bool) map[int]handStructs {
	// hand type map
	handType := map[int]handStructs{
		1: []handStruct{}, // High card
		2: []handStruct{}, // One pair
		3: []handStruct{}, // Two pair
		4: []handStruct{}, // Thre of a kind
		5: []handStruct{}, // Full House
		6: []handStruct{}, // Four of a kind
		7: []handStruct{}, // Five of a kind
	}

	// get hands into bins
	for _, line := range lines {
		lineParts := strings.Fields(line)
		//hand := {lineParts[0], util.AtoiWrapper(lineParts[1])}

		var key int
		if joker {
			key = getHandType(setWildCard(lineParts[0]))
		} else {
			key = getHandType(lineParts[0])
		}
		handType[key] = append(handType[key], handStruct{lineParts[0], util.AtoiWrapper(lineParts[1])})
	}

	return handType
}

func getHandType(hand string) int {
	cardCount := make(map[rune]int)

	for _, card := range hand {
		cardCount[card] += 1
	}

	if len(cardCount) == 5 {
		return 1 // must be high card
	} else if len(cardCount) == 4 {
		return 2 // must be one pair
	} else if len(cardCount) == 3 {
		for _, v := range cardCount {
			if v == 3 {
				return 4 // 3 different cards, and 3 of one kind -> must be three of a kind
			}
		}
		return 3 // 3 different card, but no set of 3 -> must be two pair
	} else if len(cardCount) == 2 {
		for _, v := range cardCount {
			if v == 4 {
				return 6 // four of a kind
			}
		}
		return 5 // 2 different cards, but no set of 4 -> must be full house
	} else {
		return 7 // must be five of a kind
	}
}

func setWildCard(hand string) string {
	if !strings.Contains(hand, "J") {
		return hand
	}

	// get all options for Joker
	newHandsData := []string{}
	for key := range cardStrength {
		newHandsData = append(newHandsData, strings.ReplaceAll(hand, "J", string(key))+" 0")
	}

	// group alternate hands into hand categories
	newHandTypes := groupInputIntoHandTypes(newHandsData, false)

	// get the hand in the highest category
	for i := 7; i > 0; i-- {
		if len(newHandTypes[i]) == 1 {
			return newHandTypes[i][0].hand
		} else if len(newHandTypes[i]) > 1 {
			return slices.MaxFunc(newHandTypes[i], compareHands).hand
		}
	}

	return hand // should never get this far
}

func calculateWinnings(handType map[int]handStructs) int {
	// calculate total winnings
	count := 1
	totalWinnings := 0
	for i := 1; i <= 7; i++ {
		slices.SortFunc(handType[i], compareHands)
		for _, val := range handType[i] {
			totalWinnings += val.bid * count
			count++
		}
	}

	return totalWinnings
}

func compareHands(a, b handStruct) int {

	for i := 0; i < 5; i++ {
		runeA, _ := utf8.DecodeRuneInString(a.hand[i:])
		runeB, _ := utf8.DecodeRuneInString(b.hand[i:])
		if cardStrength[runeA] < cardStrength[runeB] {
			return -1
		} else if cardStrength[runeA] > cardStrength[runeB] {
			return 1
		}
	}

	return 0
}
