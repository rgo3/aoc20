package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	max := 0
	rows := 127
	cols := 7

	var row, col int
	var ids map[int]bool
	ids = make(map[int]bool)

	for scanner.Scan() {
		instructions := scanner.Text()
		row = binFind(0, rows, instructions[:7])
		col = binFind(0, cols, instructions[7:])
		id := row*8 + col
		ids[id] = true
		if id > max {
			max = id
		}
	}
	fmt.Println("Part 1: ", max)

	for i := 0; i <= rows*8+cols; i++ {
		if ids[i] == false && ids[i-1] == true && ids[i+1] == true {
			fmt.Println("Part 2: ", i)
			break
		}
	}
}

func binFind(min int, max int, instructions string) int {
	// fmt.Println(instructions)
	for _, c := range instructions {
		diff := (max - min) + 1
		switch c {
		case 'F', 'L':
			max -= diff / 2
		case 'B', 'R':
			min += diff / 2
		}

		if min == max {
			return min
		}
	}

	return -1
}
