package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	readFile, err := os.Open("input/day6.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	cline := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			lines = append(lines, strings.TrimLeft(cline, " "))
			cline = ""
		}
		cline += (" " + line)
	}

	fmt.Println("Day 6 Part 1 Answer: ", daySix_PartOne(lines))
	fmt.Println("Day 6 Part 2 Answer: ", daySix_PartTwo(lines))
}

func daySix_PartOne(lines []string) int {
	result := 0
	for i := range lines {
		group := []string{}
		for j := range lines[i] {
			//fmt.Println(string(lines[i][j]))
			if string(lines[i][j]) != " " {
				if !slices.Contains(group, string(lines[i][j])) {
					group = append(group, string(lines[i][j]))
				}
			}
		}
		result += len(group)
	}
	return result
}

func daySix_PartTwo(lines []string) int {
	result := 0
	for i := range lines {
		group := []string{}
		for j := range lines[i] {
			if string(lines[i][j]) != " " {
				if !slices.Contains(group, string(lines[i][j])) {
					group = append(group, string(lines[i][j]))
				}
			}
		}
		people := strings.Split(lines[i], " ")
		min := len(people)
		for a := range group {
			r := 0
			for b := range people {
				if strings.Contains(people[b], string(group[a])) {
					r++
				}
			}
			if r == min {
				result++
			}
		}
	}
	return result
}
