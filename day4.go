package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	value string
	x     int
	y     int
	p     bool
	c     int
}

var grid = make(map[string]coordinate)
var count int

func main() {
	readFile, err := os.Open("day4.txt")
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
			a := makeCoordinate(x, y)
			b := coordinate{s[i], x, y, false, 0}
			grid[a] = b
			x += 1
		}
		y += 1
	}
	fmt.Println(grid)
	fmt.Println("+++++++++++++++++")
	for k, _ := range grid {
		checkXmas(up(grid[k]))
		checkXmas(down(grid[k]))
		checkXmas(left(grid[k]))
		checkXmas(right(grid[k]))
		checkXmas(upLeft(grid[k]))
		checkXmas(upRight(grid[k]))
		checkXmas(downLeft(grid[k]))
		checkXmas(downRight(grid[k]))
	}
	fmt.Println(grid)
	fmt.Println("Part 1: ", count)
}

func searchXmas() {

}

func checkXmas(z []coordinate) {
	if len(z) == 4 {
		forward := ""
		backward := ""
		for j := range 4 {
			forward += z[j].value
		}
		backward = reverse(forward)
		if forward == "XMAS" || backward == "SAMX" {
			count += 1
			for p := range 4 {
				m := makeCoordinate(z[p].x, z[p].y)
				if grid[m].p == true {
					e := coordinate{grid[m].value, grid[m].x, grid[m].y, grid[m].p, grid[m].c + 1}
					grid[m] = e
				} else {
					e := coordinate{grid[m].value, grid[m].x, grid[m].y, true, grid[m].c}
					grid[m] = e
				}
			}
		}
	}
}

func up(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x, y-i)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}

	return z
}

func down(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x, y+i)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func left(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x-i, y)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func right(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x+i, y)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func upLeft(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x-i, y-i)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func upRight(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x+i, y-i)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func downLeft(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x-i, y+i)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func downRight(c coordinate) []coordinate {
	var z []coordinate
	x := c.x
	y := c.y
	for i := range 4 {
		a := makeCoordinate(x+i, y+i)
		if value, ok := grid[a]; ok {
			z = append(z, value)
		}
	}
	return z
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func getCoordinate(a string) (int, int) {
	b := strings.Split(a, ":")
	bx, _ := strconv.Atoi(b[0])
	by, _ := strconv.Atoi(b[1])
	return bx, by
}

func makeCoordinate(x int, y int) string {
	return (strconv.Itoa(x) + ":" + strconv.Itoa(y))
}
