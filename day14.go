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
	rows int = 7
	cols int = 11
)

func main() {
	robots := []robot{}
	readFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		sx, sy := parseRobots(s[0])
		vx, vy := parseRobots(s[1])
		p := [][2]int{}
		p = append(p, [2]int{sx, sy})
		robots = append(robots, robot{sx, sy, vx, vy, p})
	}

	moveRobots(robots)
	fmt.Println(robots)

	readFile.Close()
}

func moveRobots(robots []robot) {
	for i := range len(robots) {
		updateX, updateY := 0, 0
		current := robots[i].position[:len(robots[i].position)]
		if current[0][0]+robots[i].velocityX > 0 && current[0][0]+robots[i].velocityX <= cols {
			updateX = current[0][0] + robots[i].velocityX
		} else {
			//teleport
		}
		if current[0][1]+robots[i].velocityY > 0 && current[0][1]+robots[i].velocityY <= rows {
			updateY = current[0][1] + robots[i].velocityY
		} else {
			//teleport
		}
		update := [2]int{updateX, updateY}
		robots[i].position = append(robots[i].position, update)
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
