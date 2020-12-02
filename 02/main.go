package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var pws1, pws2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}

		fields := strings.FieldsFunc(scanner.Text(), f)
		min, _ := strconv.Atoi(fields[0])
		max, _ := strconv.Atoi(fields[1])
		c := fields[2]
		pw := fields[3]

		if isPassword1(min, max, c, pw) {
			pws1++
		}

		if isPassword2(min, max, c, pw) {
			pws2++
		}
	}

	fmt.Println(pws1)
	fmt.Println(pws2)
}

func isPassword1(min, max int, c, pw string) bool {
	count := strings.Count(pw, c)
	if count >= min && count <= max {
		return true
	}
	return false
}

func isPassword2(min, max int, c, pw string) bool {
	if string(pw[min-1]) == c && string(pw[max-1]) != c {
		return true
	}
	if string(pw[min-1]) != c && string(pw[max-1]) == c {
		return true
	}
	return false
}
