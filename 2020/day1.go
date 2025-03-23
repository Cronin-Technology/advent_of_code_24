package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	expenses := []int{}
	for fileScanner.Scan() {
		num, _ := strconv.Atoi(fileScanner.Text())
		expenses = append(expenses, num)
	}
	fmt.Println("Day 1 Part 1 Answer: ", dayOne_PartOne(expenses))
	fmt.Println("Day 1 Part 2 Answer: ", dayOne_PartTwo(expenses))
}

func dayOne_PartOne(expenses []int) int {
	for x := range expenses {
		for y := range expenses {
			if expenses[x]+expenses[y] == 2020 {
				return expenses[x] * expenses[y]
			}
		}
	}
	return 0
}

func dayOne_PartTwo(expenses []int) int {
	for x := range expenses {
		for y := range expenses {
			for z := range expenses {
				if expenses[x]+expenses[y]+expenses[z] == 2020 {
					return expenses[x] * expenses[y] * expenses[z]
				}
			}
		}
	}
	return 0
}
