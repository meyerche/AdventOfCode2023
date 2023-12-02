package util

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func ReadFile(day, dataset int) []string {
	var data []string
	
	filename := "day" + strconv.Itoa(day) + "part" + strconv.Itoa(dataset) + ".txt"
	fmt.Println(filename)
	file, err := os.Open("../data/" + filename)
	
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	
	return data	
}