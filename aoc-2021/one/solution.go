package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadData(src string) int[] {
	
	var result int[]
	
	file, _ := os.Open(src)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		result = append(result, line)
	}

	return result
}

func part1(data int[]) int {
	var increment int = 0
	var previous int = 0
	for _, value := range data {
		if value > previous {
			increment++
		}
	}
	return increment - 1 
}

func part2(data int[]) int {
	var increment int = 0
	for i := 1; i+2 < len(data); i = i+1 {
		previous := data[i-1] + data[i] + data[i+1]
		current := data[i] + data[i+1] + data[i+2]
		if current > previous { increment++ }
	}
	return increment
}

func main() {
	var src string = "input.txt"
	var data int[] = loadData(src)

	var solution1 int = part1(data)
	var solution2 int = part2(data)

	fmt.Printf("Solution for part one: %v\nSolution for part two: %v\n", solution1, solution2)
}
