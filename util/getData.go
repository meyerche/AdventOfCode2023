package util

import (
	"bufio"
	"os"
	"fmt"
)

func ReadFile(day, dataset int) 	[]string {
	var data []string
	
	filename := "day" + day + "part" + dataset + ".txt"
	file, err := os.Open("../data/" + filename)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		data = data.append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	
	return data	
}