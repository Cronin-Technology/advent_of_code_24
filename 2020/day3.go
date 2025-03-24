package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	tree string
	x    int
	y    int
}

var Forest = make(map[string]Tree)

var maxX int
var minY int

func main() {
	readFile, err := os.Open("input/day3.txt")
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
			b := Tree{s[i], x, y}
			Forest[a] = b
			x += 1
		}
		maxX = x
		y -= 1
	}
	minY = y

	readFile.Close()
	printForest()
	fmt.Println("Day 3 Part One Answer: ", dayThree_PartOne(3, 1))
	fmt.Println("Day 3 Part Two Answer: ", dayThree_PartTwo())
}

func dayThree_PartOne(sX, sY int) int {
	cX, cY := 0, 0
	result := 0
	for cY > minY {
		if Forest[makePoint(cX, cY)].tree != "." {
			result += 1
		}
		if cX+sX >= maxX {
			cX = (cX + sX) - maxX
		} else {
			cX += sX
		}
		cY -= sY
	}
	return result
}

func dayThree_PartTwo() int {
	cX := []int{1, 3, 5, 7, 1}
	cY := []int{1, 1, 1, 1, 2}
	results := []int{}
	result := 1
	for i := 0; i < len(cX); i++ {
		results = append(results, dayThree_PartOne(cX[i], cY[i]))
		result = result * results[len(results)-1]
	}
	return result
}

func printForest() {
	sX, sY := 0, 0
	for sY > minY {
		con := ""
		for sX < maxX {
			con += Forest[makePoint(sX, sY)].tree
			sX += 1
		}
		sY -= 1
		sX = 0
		fmt.Println(con)
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
