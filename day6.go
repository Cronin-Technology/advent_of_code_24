package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type point struct {
	value string
	x     int
	y     int
	p     bool
	d     []string
	c     int
}

// var points = make(map[string]point)
// var paradox = make(map[string]point)
// var start []int
// var direction string
// var step int
// var maxX int
// var maxY int
// var pointCount int

func main() {
	readFile, err := os.Open("test.txt")
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
			a := makePoint(x, y)
			b := point{s[i], x, y, false, []string{}, 0}
			if s[i] == "^" {
				z, e := getPoint(a)
				start = append(start, z)
				start = append(start, e)
				direction = "up"
			}
			points[a] = b
			paradox[a] = b
			x += 1
		}
		maxX = x
		y += 1
	}
	maxY = y

	readFile.Close()
	step = 0
	GuardMovement()
	Looper()
	for i := range maxX {
		//con := ""
		for j := range maxY {
			a := makePoint(j, i)
			if points[a].value == "X" || points[a].value == "^" {
				step += 1
			}
			//con += points[a].value
		}
		//fmt.Println(con)
	}
	fmt.Println("Part 1: ", step)

	fmt.Println("Part 2: ", pointCount)
}

func GuardMovement() {
	w := 0
	next := []int{0, 0}
	direction := []int{0, -1}
	current := start
	di := append(points[makePoint(current[0], current[1])].d, getDirection(direction))
	points[makePoint(current[0], current[1])] = point{"X", points[makePoint(current[0], current[1])].x, points[makePoint(current[0], current[1])].y, true, di, points[makePoint(current[0], current[1])].c + 1}
	for ok := true; ok; ok = (w != 2) { //do while loop
		next[0] = (current[0] + direction[0])
		next[1] = (current[1] + direction[1])
		a := makePoint(next[0], next[1])
		if points[a].value == "#" {
			current[0] = (current[0] + (-1 * direction[0]))
			current[1] = (current[1] + (-1 * direction[1]))
			direction = turnRight(direction)
		} else if next[0] > maxX || next[1] > maxY || next[0] < 0 || next[1] < 0 {
			//Out of Bounds
			w = 2
		} else {
			di := append(points[a].d, getDirection(direction))
			points[a] = point{"X", points[a].x, points[a].y, true, di, points[a].c + 1}
			current = next
		}
	}
}

func Looper() {
	for i := range maxX {
		//con := ""
		for j := range maxY {
			para := make(map[string]point)
			for xx := range maxX {
				//con := ""
				for yy := range maxY {
					a := makePoint(yy, xx)
					para[a] = paradox[a]
					//con += points[a].value
				}
				//fmt.Println(con)
			}
			a := makePoint(j, i)
			if para[a].value != "^" || para[a].value != "#" {
				para[a] = point{"#", para[a].x, para[a].y, false, []string{}, para[a].c}
				if GuardMovementParadox(para) {
					pointCount += 1
				}
				para = map[string]point{}
			}
			//con += points[a].value
		}
		//fmt.Println(con)
	}
}

func GuardMovementParadox(para map[string]point) bool {
	w := 0
	l := 0
	next := []int{0, 0}
	direction := []int{0, -1}
	current := start
	di := append(para[makePoint(current[0], current[1])].d, getDirection(direction))
	para[makePoint(current[0], current[1])] = point{"X", para[makePoint(current[0], current[1])].x, para[makePoint(current[0], current[1])].y, true, di, para[makePoint(current[0], current[1])].c + 1}
	for ok := true; ok; ok = (w != 2) { //do while loop
		next[0] = (current[0] + direction[0])
		next[1] = (current[1] + direction[1])
		a := makePoint(next[0], next[1])
		if para[a].value == "#" {
			current[0] = (current[0] + (-1 * direction[0]))
			current[1] = (current[1] + (-1 * direction[1]))
			direction = turnRight(direction)
		} else if next[0] > maxX || next[1] > maxY || next[0] < 0 || next[1] < 0 {
			//Out of Bounds
			return false
		} else {
			if para[a].p {
				if slices.Contains(para[a].d, getDirection(direction)) {
					//fmt.Println(para[a].d, getDirection(direction))
					return true
				}
			}
			di := append(para[a].d, getDirection(direction))
			para[a] = point{"X", para[a].x, para[a].y, true, di, para[a].c + 1}
			current = next
			l += 1
		}
	}
	return false
}

func getDirection(d []int) string {
	if d[0] == 0 && d[1] == -1 {
		return "up"
	} else if d[0] == 1 && d[1] == 0 {
		return "right"
	} else if d[0] == 0 && d[1] == 1 {
		return "down"
	} else if d[0] == -1 && d[1] == 0 {
		return "left"
	}
	return "broken"
}

func turnRight(d []int) []int {
	if d[0] == 0 && d[1] == -1 {
		return []int{1, 0}
	} else if d[0] == 1 && d[1] == 0 {
		return []int{0, 1}
	} else if d[0] == 0 && d[1] == 1 {
		return []int{-1, 0}
	} else if d[0] == -1 && d[1] == 0 {
		return []int{0, -1}
	}
	return d
}

func getPoint(a string) (int, int) {
	b := strings.Split(a, ":")
	bx, _ := strconv.Atoi(b[0])
	by, _ := strconv.Atoi(b[1])
	return bx, by
}

func makePoint(x int, y int) string {
	return (strconv.Itoa(x) + ":" + strconv.Itoa(y))
}
