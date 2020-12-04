package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	fmt.Println(treesOnRoute(lines, 3, 1))
	fmt.Println(treesOnRoute(lines, 1, 1) * treesOnRoute(lines, 3, 1) * treesOnRoute(lines, 5, 1) * treesOnRoute(lines, 7, 1) * treesOnRoute(lines, 1, 2))

}

func treesOnRoute(grid []string, right, down int) int {
	position := right
	trees := 0

	gridWidth := len(grid[0])

	for i := down; i < len(grid); i += down {
		if string(grid[i][position%gridWidth]) == "#" {
			trees++
		}
		position += right
	}
	return trees
}
