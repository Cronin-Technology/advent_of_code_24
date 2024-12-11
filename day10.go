package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ROWS, COLS int

var directions = []struct{ dx, dy int }{
	{1, 0},  // down
	{-1, 0}, // up
	{0, 1},  // right
	{0, -1}, // left
}

func bfs(mapData [][]int, startX, startY, targetX, targetY int) (int, [][][2]int) {
	queue := [][]int{{startX, startY}} // BFS queue (stores coordinates)
	paths := map[[2]int][][][2]int{}   // To store all paths ending at a position
	paths[[2]int{startX, startY}] = [][][2]int{{{startX, startY}}}
	visited := make(map[[2]int]bool)

	for len(queue) > 0 {
		// Get the current position
		curr := queue[0]
		queue = queue[1:]
		x, y := curr[0], curr[1]

		// Check if we reached the target (elevation 9)
		if x == targetX && y == targetY {
			// Collect all paths that lead to the target
			var result [][][2]int
			for _, path := range paths[[2]int{x, y}] {
				result = append(result, path)
			}
			return len(result), result
		}

		// Explore neighboring cells
		for _, dir := range directions {
			newX, newY := x+dir.dx, y+dir.dy
			if newX >= 0 && newX < ROWS && newY >= 0 && newY < COLS && !visited[[2]int{newX, newY}] && mapData[newX][newY] == mapData[x][y]+1 {
				visited[[2]int{newX, newY}] = true
				// Add to the queue and update paths
				queue = append(queue, []int{newX, newY})
				for _, path := range paths[[2]int{x, y}] {
					newPath := append([][2]int{}, path...) // Make a copy of the path
					newPath = append(newPath, [2]int{newX, newY})
					paths[[2]int{newX, newY}] = append(paths[[2]int{newX, newY}], newPath)
				}
			}
		}
	}
	return 0, nil
}

func main() {
	count := 0
	// mapData := [][]int{
	// 	{0, 1, 2, 3},
	// 	{1, 2, 3, 4},
	// 	{8, 7, 6, 5},
	// 	{9, 8, 7, 6},
	// }
	mapData := [][]int{}
	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		input := strings.Split(fileScanner.Text(), "")
		temp := []int{}
		for i := range len(input) {
			a, _ := strconv.Atoi(input[i])
			temp = append(temp, a)
		}
		ROWS = len(input)
		COLS = len(input)
		mapData = append(mapData, temp)
	}

	for k := 0; k < ROWS; k++ {
		for l := 0; l < COLS; l++ {
			if mapData[k][l] == 9 {
				targetX, targetY := l, k
				for i := 0; i < ROWS; i++ {
					for j := 0; j < COLS; j++ {
						if mapData[i][j] == 0 {
							startX, startY := j, i

							_, paths := bfs(mapData, startX, startY, targetX, targetY)
							if paths != nil {
								fmt.Println("Paths from 0 to 9 with increments of 1:")
								// for _, path := range paths {
								// 	fmt.Println(len(path))
								// }
								count += 1
							} else {
								fmt.Println("No paths found.")
							}
						}
					}
				}
			}
		}
	}
	// Start BFS from position (0, 0) (elevation 0)

	fmt.Println("Part 1:", count)
}
