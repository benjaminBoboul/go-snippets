package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func part1() Position {
	var pos Position = Position{0, 0}
	var rx *regexp.Regexp = regexp.MustCompile(`(\w+) (\d+)`)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := rx.FindAllStringSubmatch(line, -1)

		v, _ := strconv.Atoi(matches[0][2])

		switch matches[0][1] {
		case "down":
			pos.Y += v
		case "up":
			pos.Y -= v
		case "forward":
			pos.X += v
		default:
			panic("Strange direction capt'n " + matches[0][1])
		}
	}

	return pos
}

func part2() Position {
	var pos Position = Position{0, 0}
	var rx *regexp.Regexp = regexp.MustCompile(`(\w+) (\d+)`)
	var aim int = 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := rx.FindAllStringSubmatch(line, -1)

		v, _ := strconv.Atoi(matches[0][2])

		switch matches[0][1] {
		case "down":
			aim += v
		case "up":
			aim -= v
		case "forward":
			pos.X += v
			pos.Y += v * aim
		default:
			panic("Strange direction capt'n " + matches[0][1])
		}
	}

	return pos
}

func main() {
	result := part1()

	fmt.Println(result.X * result.Y)

	result = part2()

	fmt.Println(result.X * result.Y)
}
