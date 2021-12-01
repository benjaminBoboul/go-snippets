package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countIncrement(src string) int {
	file, _ := os.Open(src)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	incrementCounter := 0
	previousIter := 0

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		if line > previousIter {
			incrementCounter += 1	
		}

		previousIter = line
	}
	return incrementCounter - 1 // remove one increment caused by previousIter starting at 0 
}

func main() {
	var src string = "input.txt"
	var result int = countIncrement(src)

	fmt.Printf("I've incremented %v times", result)
}
