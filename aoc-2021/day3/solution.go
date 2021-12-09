package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
)

func loadData(src string) [1000][12]int {

	var result [1000][12]int

	file, _ := os.Open(src)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var i int = 0
	for scanner.Scan() {
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			value, _ := strconv.Atoi(string(line[j]))
			result[i][j] = value
		}
		i++
	}

	return result
}

func part1 () {
	data := loadData("input.txt")
	var nb_col int = len(data[0])
	var nb_line int = len(data)

	var decoded_gamma []int

	for i := 0; i < nb_col; i++ {
		var tmp_sum int = 0
		for j := 0; j < nb_line; j++ {
			tmp_sum += data[j][i]
		}
		fmt.Println(nb_line)
		fmt.Println(tmp_sum)
		if (tmp_sum >= nb_line/2) {
			decoded_gamma = append(decoded_gamma, 1)
		} else {
			decoded_gamma = append(decoded_gamma, 0)
		}
	}

	fmt.Println(decoded_gamma)
}

func part2(data [][]int) {


func main () {
	part1()
}
