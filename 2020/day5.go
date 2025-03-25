package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	readFile, err := os.Open("input/day5.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	seats := []string{}
	for fileScanner.Scan() {
		seats = append(seats, fileScanner.Text())
	}
	fmt.Println("Day 5 Part 1 Answer: ", dayFive_PartOne(seats))
	fmt.Println("Day 5 Part 2 Answer: ", dayFive_PartTwo(seats))
}

func dayFive_PartOne(seats []string) int {
	result := 0
	resultsRow := []int{}
	resultsCol := []int{}

	for i := range seats {
		max := 127
		colmax := 7
		cnt := 128
		colcnt := 8
		min := 0
		colmin := 0
		for j := range seats[i] {
			if j < 7 {
				cnt = cnt / 2
				// fmt.Println(string(seats[i][j]))
				// fmt.Println("Total Numbers", cnt)
				if string(seats[i][j]) == "F" {
					max = min + (cnt - 1)
					if cnt == 1 {
						resultsRow = append(resultsRow, min)
					}
				} else {
					min = max - (cnt - 1)
					if cnt == 1 {
						resultsRow = append(resultsRow, max)
					}
				}
			}
			if j >= 7 {
				colcnt = colcnt / 2
				// fmt.Println(string(seats[i][j]))
				// fmt.Println("Total Numbers", cnt)
				if string(seats[i][j]) == "L" {
					colmax = colmin + (colcnt - 1)
					if colcnt == 1 {
						resultsCol = append(resultsCol, colmin)
					}
				} else {
					colmin = colmax - (colcnt - 1)
					if colcnt == 1 {
						resultsCol = append(resultsCol, colmax)
					}
				}
			}
		}
	}
	//fmt.Println(resultsRow, resultsCol)
	top := 0
	for f := range resultsRow {
		tmp := (resultsRow[f] * 8) + resultsCol[f]
		if tmp > top {
			top = tmp
		}
	}
	result = top
	return result
}

func dayFive_PartTwo(seats []string) int {
	resultsRow := []int{}
	resultsCol := []int{}

	for i := range seats {
		max := 127
		colmax := 7
		cnt := 128
		colcnt := 8
		min := 0
		colmin := 0
		for j := range seats[i] {
			if j < 7 {
				cnt = cnt / 2
				// fmt.Println(string(seats[i][j]))
				// fmt.Println("Total Numbers", cnt)
				if string(seats[i][j]) == "F" {
					max = min + (cnt - 1)
					if cnt == 1 {
						resultsRow = append(resultsRow, min)
					}
				} else {
					min = max - (cnt - 1)
					if cnt == 1 {
						resultsRow = append(resultsRow, max)
					}
				}
			}
			if j >= 7 {
				colcnt = colcnt / 2
				// fmt.Println(string(seats[i][j]))
				// fmt.Println("Total Numbers", cnt)
				if string(seats[i][j]) == "L" {
					colmax = colmin + (colcnt - 1)
					if colcnt == 1 {
						resultsCol = append(resultsCol, colmin)
					}
				} else {
					colmin = colmax - (colcnt - 1)
					if colcnt == 1 {
						resultsCol = append(resultsCol, colmax)
					}
				}
			}
		}
	}
	//fmt.Println(resultsRow, resultsCol)
	top := 0
	seatids := []int{}
	for f := range resultsRow {
		tmp := (resultsRow[f] * 8) + resultsCol[f]
		if tmp > top {
			top = tmp
		}
		seatids = append(seatids, tmp)
	}
	sort.Sort(sort.IntSlice(seatids))
	for q := seatids[0]; q < len(seatids); q++ {
		if !slices.Contains(seatids, q) {
			return q
		}
	}
	return 0
}
