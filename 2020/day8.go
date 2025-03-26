package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input/day8.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	instructions := [][]string{}

	for fileScanner.Scan() {
		instruction := []string{}
		a := strings.Split(fileScanner.Text(), " ")
		instruction = append(instruction, a[0])
		instruction = append(instruction, a[1])
		instructions = append(instructions, instruction)
	}
	fmt.Println("Day 8 Part 1 Answer: ", dayEight_PartOne(instructions))
	fmt.Println("Day 8 Part 2 Answer: ", dayEight_PartTwo(instructions))
}

func dayEight_PartOne(i [][]string) int {
	result := 0
	visited := []int{}
	place := 0
	repeat := false
	for !repeat {
		//fmt.Println(i[place][0], i[place][1])
		if slices.Contains(visited, place) {
			repeat = true
		} else {
			if i[place][0] == "nop" {
				visited = append(visited, place)
				place++
			} else if i[place][0] == "acc" {
				visited = append(visited, place)
				z, _ := strconv.Atoi(i[place][1])
				result += z
				place++
			} else if i[place][0] == "jmp" {
				visited = append(visited, place)
				z, _ := strconv.Atoi(i[place][1])
				place += z
			}
		}

	}
	return result
}

func dayEight_PartTwo(i [][]string) int {
	complete := false
	changed := []int{}
	iter := 0
	for !complete {
		changed = append(changed, iter)
		result := 0
		visited := []int{}
		place := 0
		repeat := false
		for !repeat {
			//fmt.Println(i[place][0], i[place][1])
			if slices.Contains(visited, place) {
				repeat = true
			} else {
				if i[place][0] == "nop" && place != iter {
					visited = append(visited, place)
					place++
				} else if i[place][0] == "acc" {
					visited = append(visited, place)
					z, _ := strconv.Atoi(i[place][1])
					result += z
					place++
				} else if i[place][0] == "jmp" && place != iter {
					visited = append(visited, place)
					z, _ := strconv.Atoi(i[place][1])
					place += z
				} else if i[place][0] == "jmp" && place == iter {
					visited = append(visited, place)
					place++
				} else if i[place][0] == "nop" && place == iter {
					visited = append(visited, place)
					z, _ := strconv.Atoi(i[place][1])
					place += z
				}
			}
			if place >= len(i) {
				repeat = true
				complete = true
				return result
			}
		}
		iter++
	}

	return 0
}
