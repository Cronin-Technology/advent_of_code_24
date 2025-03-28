package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Seat struct {
	v    string
	x    int
	y    int
	flip bool
}

var Seating = [][2]int{}
var Seats = []Seat{}

var maxSeatX int
var minSeatY int

var PreviousSeating []string

func main() {
	readFile, err := os.Open("input/day11.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	y := 0
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), "")
		x := 0
		for i := range len(s) {
			//a := makePoint(x, y)
			b := Seat{s[i], x, y, false}
			Seating = append(Seating, [2]int{x, y})
			Seats = append(Seats, b)
			x += 1
		}
		maxSeatX = x
		y -= 1
	}
	minSeatY = y
	printSeating()
	fmt.Println("____________________________")
	readFile.Close()
	fmt.Println("Day 11 Part One Answer: ", dayEleven_PartOne())
	fmt.Println("Day 11 Part Two Answer: ", dayEleven_PartTwo())
}

func dayEleven_PartOne() int {
	result := 0
	sX, sY := 0, 0
	for sY > minSeatY {
		for sX < maxSeatX {
			a := slices.Index(Seating, [2]int{sX, sY})
			PreviousSeating = append(PreviousSeating, Seats[a].v)
			sX += 1
		}
		sY -= 1
		sX = 0
	}

	findFlips()
	executeFlips()
	for !checkChange() {
		sX, sY = 0, 0
		PreviousSeating = []string{}
		for sY > minSeatY {
			for sX < maxSeatX {
				a := slices.Index(Seating, [2]int{sX, sY})
				PreviousSeating = append(PreviousSeating, Seats[a].v)
				sX += 1
			}
			sY -= 1
			sX = 0
		}
		findFlips()
		executeFlips()
	}
	sX, sY = 0, 0
	for sY > minSeatY {
		for sX < maxSeatX {
			a := slices.Index(Seating, [2]int{sX, sY})
			if Seats[a].v == "#" {
				result++
			}
			sX += 1
		}
		sY -= 1
		sX = 0
	}
	printSeating()
	return result
}

func findFlips() {
	sX, sY := 0, 0
	for sY > minSeatY {
		for sX < maxSeatX {
			update := false
			occupied := 0
			empty := 0
			tmp := Seats[slices.Index(Seating, [2]int{sX, sY})].v
			adj := getAdjacent([2]int{sX, sY})
			for a := range adj {
				if slices.Index(Seating, adj[a]) != -1 {
					if Seats[slices.Index(Seating, adj[a])].v == "#" {
						occupied++
					} else if Seats[slices.Index(Seating, adj[a])].v == "L" {
						empty++
					}
				}
			}
			if tmp == "L" && occupied == 0 {
				update = true
			} else if tmp == "#" && occupied >= 4 {
				update = true
			} else if tmp == "." {
				update = false
			} else {
				update = false
			}
			Seats[slices.Index(Seating, [2]int{sX, sY})].flip = update
			sX += 1
		}
		sY -= 1
		sX = 0
	}
}

func executeFlips() {
	sX, sY := 0, 0
	for sY > minSeatY {
		for sX < maxSeatX {
			update := Seats[slices.Index(Seating, [2]int{sX, sY})].v
			if Seats[slices.Index(Seating, [2]int{sX, sY})].v == "L" && Seats[slices.Index(Seating, [2]int{sX, sY})].flip {
				update = "#"
			} else if Seats[slices.Index(Seating, [2]int{sX, sY})].v == "#" && Seats[slices.Index(Seating, [2]int{sX, sY})].flip {
				update = "L"
			}
			Seats[slices.Index(Seating, [2]int{sX, sY})].v = update
			sX += 1
		}
		sY -= 1
		sX = 0
	}
}

func checkChange() bool {
	sX, sY := 0, 0
	currentSeating := []string{}
	for sY > minSeatY {
		for sX < maxSeatX {
			currentSeating = append(currentSeating, Seats[slices.Index(Seating, [2]int{sX, sY})].v)
			sX += 1
		}
		sY -= 1
		sX = 0
	}
	if len(PreviousSeating) != len(currentSeating) {
		return false
	}
	for i := range PreviousSeating {
		if PreviousSeating[i] != currentSeating[i] {
			return false
		}
	}
	return true
}

func dayEleven_PartTwo() int {
	return 0
}

func printSeating() {
	sX, sY := 0, 0
	for sY > minSeatY {
		con := ""
		for sX < maxSeatX {
			con += Seats[slices.Index(Seating, [2]int{sX, sY})].v
			sX += 1
		}
		sY -= 1
		sX = 0
		fmt.Println(con)
	}
}

func getAdjacent(start [2]int) [][2]int {
	result := [][2]int{}
	a, b := start[0], start[1]
	result = append(result, [2]int{a + 1, b})
	result = append(result, [2]int{a + 1, b - 1})
	result = append(result, [2]int{a - 1, b + 1})
	result = append(result, [2]int{a - 1, b})
	result = append(result, [2]int{a, b + 1})
	result = append(result, [2]int{a, b - 1})
	result = append(result, [2]int{a - 1, b - 1})
	result = append(result, [2]int{a + 1, b + 1})
	return result
}
