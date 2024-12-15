package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	value    string
	x        int
	y        int
	occupied bool
}

type bot struct {
	startX, startY     int
	currentX, currentY int
}

var warehouse = make(map[string]tile)
var walle bot
var instructions []string
var maxCols, maxRows int

func main() {
	readFile, err := os.Open("day15.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	row := 0
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), "")
		col := 0
		for i := range len(s) {
			a := makePoint(col, row)
			b := tile{s[i], col, row, false}
			if s[i] == "@" {
				walle = bot{col, row, col, row}
				b = tile{".", col, row, false}
			}
			warehouse[a] = b
			col += 1
		}
		maxCols = col
		row += 1
	}
	maxRows = row

	readFile.Close()
	readFile1, err := os.Open("day15_2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner1 := bufio.NewScanner(readFile1)
	fileScanner1.Split(bufio.ScanLines)
	for fileScanner1.Scan() {
		s := strings.Split(fileScanner1.Text(), "")
		for i := range len(s) {
			instructions = append(instructions, s[i])
		}
	}

	printWarehouse()

	for i := range len(instructions) {
		moveBot(instructions[i])
		//printWarehouse()
	}

	printWarehouse()

	fmt.Println("Part 1 : ", calculateWarehouse())

}

func calculateWarehouse() int {
	count := 0
	for i := range maxRows {
		for j := range maxCols {
			if warehouse[makePoint(j, i)].value == "O" {
				count += (i * 100) + j
			}
		}
	}
	return count
}

func moveBot(move string) {
	walleLoc := makePoint(walle.currentX, walle.currentY)
	next := getMoveDirection(walleLoc, move)
	if warehouse[next].value == "#" {
		//Do Nothing - no movement or boxes to evaluate
	} else if warehouse[next].value == "." {
		//Just move the bot without any box interactions
		a, b := getPoint(next)
		walleUpdate := bot{walle.startX, walle.startY, a, b}
		walle = walleUpdate
	} else if warehouse[next].value == "O" {
		toMove := getEndofBoxes(next, move)
		if toMove {
			a, b := getPoint(next)
			walleUpdate := bot{walle.startX, walle.startY, a, b}
			walle = walleUpdate
		} else {
			//Do Nothing
		}
	} else {
		//Do Nothing
	}
}

func getEndofBoxes(c, s string) bool {
	r := 0
	toUpdate := []string{}
	toUpdate = append(toUpdate, c)
	for ok := true; ok; ok = (r != 1) {
		next := getMoveDirection(toUpdate[len(toUpdate)-1 : len(toUpdate)][0], s)
		if warehouse[next].value == "O" {
			toUpdate = append(toUpdate, next)
		} else if warehouse[next].value == "#" {
			//fmt.Println("End of this")
			return false
		} else if warehouse[next].value == "." {
			toUpdate = append(toUpdate, next)
			r = 1
		}
	}
	if len(toUpdate) > 1 {
		front := toUpdate[:1][0]
		fx, fy := getPoint(front)
		back := toUpdate[len(toUpdate)-1 : len(toUpdate)][0]
		bx, by := getPoint(back)
		f := tile{".", fx, fy, true}
		b := tile{"O", bx, by, true}
		warehouse[front] = f
		warehouse[back] = b
		return true
	}
	return false
}

func getMoveDirection(c, s string) string {
	a, b := getPoint(c)
	if s == "<" {
		return makePoint(a-1, b)
	}
	if s == "^" {
		return makePoint(a, b-1)
	}
	if s == ">" {
		return makePoint(a+1, b)
	}
	if s == "v" {
		return makePoint(a, b+1)
	}
	return makePoint(a, b)
}

func printWarehouse() {
	for i := range maxRows {
		con := ""
		for j := range maxCols {
			if walle.currentX == j && walle.currentY == i {
				con += "@"
			} else {
				con += (warehouse[makePoint(j, i)].value)
			}
		}
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
