package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	value string
	x     int
	y     int
	p     bool
	c     int
}

var points = make(map[string]point)
var start string
var direction string
var step int
var maxX int
var maxY int

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
			b := point{s[i], x, y, false, 0}
			if s[i] == "^" {
				start = a
				direction = "up"
			}
			points[a] = b
			x += 1
		}
		maxX = x
		y += 1
	}
	maxY = y

	readFile.Close()
	step = 0
	fmt.Println(points)
}

func GuardMovement() {
	w := 0
	if w != 1 { //do while loop

	}
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
