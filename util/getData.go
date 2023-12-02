package util

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func ReadFile(day, dataset int) []string {
	var data []string
	
	folder := "day" + strconv.Itoa(day) + "/"
	filename := "part" + strconv.Itoa(dataset) + ".txt"

		file, err := os.Open("../data/" + folder + filename)
	
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