package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	diff := 3
	adapters := []int{}
	adaptersList := []int{}
	readFile, err := os.Open("input/day10.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		a, _ := strconv.Atoi(fileScanner.Text())
		adapters = append(adapters, a)
		adaptersList = append(adaptersList, a)
	}
	fmt.Println("Day 10 Part 1 Answer: ", dayTen_PartOne(adapters, diff))
	fmt.Println("Day 10 Part 2 Answer: ", dayTen_PartTwo(adaptersList, diff))
}

func dayTen_PartOne(adapters []int, diff int) int {
	one := 0
	three := 0
	adapters = append(adapters, 0)
	slices.Sort(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)
	for a := range adapters {
		if a+1 >= len(adapters) {
			break
		}
		if adapters[a+1]-adapters[a] == 1 {
			one++
		} else if adapters[a+1]-adapters[a] == 3 {
			three++
		}
	}
	fmt.Println(one, three)
	return (one * three)
}

func dayTen_PartTwo(adapters []int, diff int) int {
	result := 1
	cnt := 0
	one := 0
	three := 0
	adapters = append(adapters, 0)
	slices.Sort(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)
	blocks := [][]int{}
	cblock := []int{0}
	for a := range adapters {
		if a+1 >= len(adapters) {
			blocks = append(blocks, []int{adapters[a]})
			break
		}
		if adapters[a+1]-adapters[a] == 1 {
			one++
			cblock = append(cblock, adapters[a+1])
		} else if adapters[a+1]-adapters[a] == 3 {
			three++
			blocks = append(blocks, cblock)
			cblock = []int{adapters[a+1]}
		}
	}
	for b := range blocks {
		if b+1 >= len(blocks) {
			break
		}
		if len(blocks[b]) == 4 {
			result *= 4
			cnt++
		}
		if len(blocks[b]) == 3 {
			result *= 2
			cnt++
		}
		if len(blocks[b]) == 5 {
			result *= 7
			cnt++
		}
	}
	fmt.Println(one, three)
	fmt.Println(result, cnt)
	fmt.Println(blocks)
	return result
}
