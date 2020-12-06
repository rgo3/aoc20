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

	var member int
	var sum1, sum2 int
	answers1 := make(map[rune]bool)
	answers2 := make(map[rune]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			member++
			for _, c := range line {
				answers2[c]++
				if _, ok := answers1[c]; !ok {
					answers1[c] = true
				}
			}
			continue
		}

		sum1 += len(answers1)

		for _, v := range answers2 {
			if member == v {
				sum2++
			}
		}

		answers1 = make(map[rune]bool)
		answers2 = make(map[rune]int)
		member = 0
	}
	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}
