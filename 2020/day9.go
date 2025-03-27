package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	preamble := 25
	xmas := []int{}
	readFile, err := os.Open("input/day9.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		a, _ := strconv.Atoi(fileScanner.Text())
		xmas = append(xmas, a)
	}
	val := dayNine_PartOne(xmas, preamble)
	fmt.Println("Day 9 Part 1 Answer: ", val)
	fmt.Println("Day 9 Part 2 Answer: ", dayNine_PartTwo(xmas, val))
}

func dayNine_PartOne(xmas []int, preamble int) int {
	result := 0
	for i := preamble; i < len(xmas)-1; i++ {
		check := xmas[i-preamble : i]
		value := xmas[i]
		valid := 0
		for j := range check {
			if slices.Contains(check, value-check[j]) {
				valid++
			}
		}
		if valid == 0 {
			i = len(xmas)
			return value
		}
	}
	return result
}

func dayNine_PartTwo(xmas []int, val int) int {
	result := 0
	check := []int{}
	i := 0
	for len(check) <= len(xmas) {
		j := i
		for sumSlice(check) <= val {
			if sumSlice(check) == val {
				slices.Sort(check)
				return (check[0] + check[len(check)-1])
			}
			check = append(check, xmas[j])
			j++
		}
		check = []int{}
		i++
	}

	return result
}

func sumSlice(a []int) int {
	result := 0
	for i := range a {
		result += a[i]
	}
	return result
}
