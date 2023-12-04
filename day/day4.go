package day

import (
	"fmt"
	"slices"
	"strings"
)

func Day4Part1(cards []string) int {
	fmt.Println(" -- Starting part 1 --")

	sum := 0

	for _, card := range cards {
		winningNums, selectedNums := splitCard(card)

		// fmt.Println(winningNums)
		// fmt.Println(selectedNums)

		n := numberOfWins(winningNums, selectedNums)
		// fmt.Println(n)

		sum += calculatePoints(n)
	}

	return sum
}

func Day4Part2(cards []string) int {
	fmt.Println(" -- Starting part 2 --")

	cardCounts := make([]int, len(cards)) // initializes all card counts to 0

	for i, card := range cards {
		cardCounts[i] += 1 //account for original card

		winningNums, selectedNums := splitCard(card)

		n := numberOfWins(winningNums, selectedNums)

		for j := 1; j <= n; j++ {
			cardCounts[i+j] += cardCounts[i]
		}
	}

	return countAllCards(cardCounts)
}

func splitCard(card string) ([]string, []string) {
	f := func(c rune) bool {
		return c == ':' || c == '|'
	}

	// split into card number, winning numbers, and selected numbers
	cardParts := strings.FieldsFunc(card, f)

	return strings.Fields(cardParts[1]), strings.Fields(cardParts[2])
}

func numberOfWins(winningNums, selectedNums []string) int {
	delFunc := func(num string) bool {
		return !slices.Contains(winningNums, num)
	}

	wins := slices.DeleteFunc(selectedNums, delFunc)

	return len(wins)
}

func calculatePoints(n int) int {
	if n == 0 {
		return 0
	} else {
		return 1 << (n - 1)
	}
}

func countAllCards(counts []int) int {
	sum := 0
	for _, count := range counts {
		sum += count
	}
	return sum
}
