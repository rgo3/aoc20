package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, expense)
	}

	em := make(map[int]int)

	for _, e := range expenses {
		em[2020-e] = e
	}

	for _, e := range expenses {
		if v, ok := em[e]; ok {
			fmt.Println(v * e)
			break
		}
	}
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, expense)
	}

	em := make(map[int]int)

	for _, e1 := range expenses {
		for _, e2 := range expenses {
			if e2 == e1 {
				continue
			}
			em[2020 - e1 - e2] = e1 * e2
		}
	}

	for _, e := range expenses {
		if v, ok := em[e]; ok {
			fmt.Println(v * e)
			break
		}
	}
}
