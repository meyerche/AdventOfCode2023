package util

import "strconv"

func AtoiWrapper(s string) int {
	// I wouldn't normally ignore potential errors like this.
	n, _ := strconv.Atoi(s)
	return n
}
