package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	hgt = regexp.MustCompile(`(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)`)
	hcl = regexp.MustCompile(`#[0-9a-f]{6}`)
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var valid1, valid2 int
	passport := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			pairs := strings.Split(line, " ")
			for _, p := range pairs {
				kv := strings.Split(p, ":")
				passport[kv[0]] = kv[1]
			}
			continue
		}

		if isValidPassport1(passport) {
			valid1++
		}
		if isValidPassport2(passport) {
			valid2++
		}
		passport = make(map[string]string)
	}
	fmt.Println("Part 1: ", valid1)
	fmt.Println("Part 2: ", valid2)
}

func isValidPassport1(p map[string]string) bool {
	if len(p) > 7 {
		return true
	}

	if len(p) == 7 {
		if _, ok := p["cid"]; !ok {
			return true
		}
	}
	return false
}

func isValidPassport2(p map[string]string) bool {
	if len(p) < 7 {
		return false
	}

	if len(p) == 7 {
		if _, ok := p["cid"]; ok {
			return false
		}
	}

	for k, v := range p {
		switch k {
		case "byr":
			y, _ := strconv.Atoi(v)
			if y < 1920 || y > 2002 {
				return false
			}
		case "iyr":
			y, _ := strconv.Atoi(v)
			if y < 2010 || y > 2020 {
				return false
			}
		case "eyr":
			y, _ := strconv.Atoi(v)
			if y < 2020 || y > 2030 {
				return false
			}
		case "hgt":
			if hgt.MatchString(v) == false {
				return false
			}
		case "hcl":
			if hcl.MatchString(v) == false {
				return false
			}
		case "ecl":
			if !isValidEyecolor(v) {
				return false
			}
		case "pid":
			if len(v) < 9 || len(v) > 9 {
				return false
			}

			for _, c := range v {
				if c < '0' || c > '9' {
					return false
				}
			}
		}
	}

	return true
}

func isValidEyecolor(ecl string) bool {
	switch ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}
