package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input/day13.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	start := 0
	buses := []int{}
	r := 0
	for fileScanner.Scan() {
		if r == 0 {
			a, _ := strconv.Atoi(fileScanner.Text())
			start = a
		} else {
			a := strings.Split(fileScanner.Text(), ",")
			for i := range a {
				b, err := strconv.Atoi(a[i])
				if err != nil {
					buses = append(buses, 0)
				} else {
					buses = append(buses, b)
				}
			}
		}
		r++
	}
	//slices.Sort(buses)
	//fmt.Println(start, buses)
	fmt.Println("Day 13 Part 1 Answer: ", dayThirteen_PartOne(start, buses))
	fmt.Println("Day 13 Part 2 Answer: ", dayThirteen_PartTwo(start, buses))
}

func dayThirteen_PartOne(start int, buses []int) int {
	result := 0
	place := 0
	max := buses[len(buses)-1] + start
	for a := range buses {
		//generate list to
		if buses[a] != 0 {
			i := 0
			for i <= max {
				i += buses[a]
				if i >= start && i < max {
					if place != 0 && i < place {
						place = i
						result = buses[a]
					} else if place == 0 {
						place = i
						result = buses[a]
					}
				}
			}
		}
	}
	return (place - start) * result
}

func dayThirteen_PartTwo(start int, buses []int) int {
	result := 0
	ok := false
	for !ok {

		//ok is true when slice of current matches slice of previous
	}
	return result
}
