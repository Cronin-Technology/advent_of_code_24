package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tile struct {
	x        int
	y        int
	occupied bool
	robots   []string
}

type robot struct {
	startX, startY       int
	velocityX, velocityY int
	position             [][2]int
}

const (
	rows int = 103
	cols int = 101
)

var robots = make(map[int]robot)

func main() {
	readFile, err := os.Open("day14.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	id := 0
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		sx, sy := parseRobots(s[0])
		vx, vy := parseRobots(s[1])
		p := [][2]int{}
		p = append(p, [2]int{sx, sy})
		//robots = append(robots, robot{sx, sy, vx, vy, p})
		robots[id] = robot{sx, sy, vx, vy, p}
		id += 1
	}

	for i := 0; i < 101; i++ {
		moveRobots()
	}
	//fmt.Println(robots[1])
	fmt.Println(countQuadrants())

	readFile.Close()
}

func countQuadrants() int {
	temp := [][2]int{}
	for i := range len(robots) {
		current := robots[i].position[len(robots[i].position)-1 : len(robots[i].position)]
		update := current[0]
		temp = append(temp, update)
	}
	//fmt.Println(temp)
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for i := range len(temp) {
		ignoreX := cols / 2
		ignoreY := rows / 2
		if temp[i][0] < ignoreX && temp[i][1] < ignoreY {
			q1 += 1
		} else if temp[i][0] > ignoreX && temp[i][1] < ignoreY {
			q2 += 1
		} else if temp[i][0] < ignoreX && temp[i][1] > ignoreY {
			q3 += 1
		} else if temp[i][0] > ignoreX && temp[i][1] > ignoreY {
			q4 += 1
		}
	}
	//return []int{q1, q2, q3, q4}
	fmt.Println(q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func moveRobots() {
	for i := range len(robots) {
		updateX, updateY := 0, 0
		current := robots[i].position[len(robots[i].position)-1 : len(robots[i].position)]
		//fmt.Println("current : ", current)

		if current[0][0]+robots[i].velocityX >= 0 && current[0][0]+robots[i].velocityX < cols {
			updateX = current[0][0] + robots[i].velocityX
		} else {
			if current[0][0]+robots[i].velocityX < 0 {
				updateX = cols + (current[0][0] + robots[i].velocityX)
			} else {
				updateX = (current[0][0] + robots[i].velocityX) - cols
			}
		}
		if current[0][1]+robots[i].velocityY >= 0 && current[0][1]+robots[i].velocityY < rows {
			updateY = current[0][1] + robots[i].velocityY
		} else {
			if current[0][1]+robots[i].velocityY < 0 {
				updateY = rows + (current[0][1] + robots[i].velocityY)
			} else {
				updateY = (current[0][1] + robots[i].velocityY) - rows
			}
		}
		update := [2]int{updateX, updateY}
		temp := robots[i].position
		temp = append(temp, update)
		robots[i] = robot{robots[i].startX, robots[i].startY, robots[i].velocityX, robots[i].velocityY, temp}
	}
}

func parseRobots(input string) (int, int) {
	// Define a regular expression pattern to match the values
	re := regexp.MustCompile(`([a-zA-Z])=([-]?\d+),([-]?\d+)`)

	// Find all matches using the regex
	matches := re.FindAllStringSubmatch(input, -1)

	// Loop through the matches and format the output
	for _, match := range matches {
		x, _ := strconv.Atoi(match[2])
		y, _ := strconv.Atoi(match[3])
		return x, y
	}
	return 0, 0
}
