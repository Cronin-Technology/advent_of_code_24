package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type loci struct {
	value   string
	x, y, c int
}

var maxX int
var maxY int
var v []string
var per []int

func main() {
	graph := make(map[string]loci)
	part1 := 0
	// part2 := 0
	readFile, err := os.Open("day12.txt")
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
			b := loci{s[i], x, y, 0}
			graph[a] = b
			x += 1
		}
		maxX = x
		y += 1
	}
	maxY = y

	readFile.Close()

	for i := range maxX {
		con := ""
		for j := range maxY {
			a := makePoint(j, i)
			con += graph[a].value
		}
		fmt.Println(con)
	}

	gardens := make(map[string][]string)
	gardensWalked := []string{}

	for i := range maxX {
		for j := range maxY {
			a := makePoint(j, i)
			if !slices.Contains(gardensWalked, a) {
				visited := []string{}
				gardenDFS(graph, a, visited, 0)
				gardens[a] = v
				sum := 0
				for e := range len(per) {
					sum += per[e]
				}
				part1 += sum * len(v)
				for k := range len(v) {
					gardensWalked = append(gardensWalked, v[k])
				}
				v = []string{}
				per = []int{}
			}
		}
	}

	// for k, v := range gardens{
	// 	count += getFence(v)
	// }
	fmt.Println("Part 1: ", part1)
}

// func getFence(graph map[string]loci, garden []string) int {
// 	for
// }

func gardenDFS(graph map[string]loci, vertex string, visited []string, p int) {
	n, p := getNeighbors(graph, vertex)
	visited = append(visited, vertex)
	if !slices.Contains(v, vertex) {
		per = append(per, p)
		v = append(v, vertex)
	}

	for i := range len(n) {
		if !slices.Contains(visited, n[i]) {
			gardenDFS(graph, n[i], visited, 0)
		}
	}
}

func getNeighbors(graph map[string]loci, vertex string) ([]string, int) {
	x, y := getPoint(vertex)
	p := 0
	n := []string{}
	if (x-1) > -1 && (x-1) < maxX {
		if graph[makePoint(x-1, y)].value == graph[vertex].value {
			n = append(n, makePoint(x-1, y))
		} else {
			p += 1
		}
	} else {
		p += 1
	}
	if (x+1) > -1 && (x+1) < maxX {
		if graph[makePoint(x+1, y)].value == graph[vertex].value {
			n = append(n, makePoint(x+1, y))
		} else {
			p += 1
		}
	} else {
		p += 1
	}
	if (y-1) > -1 && (y-1) < maxX {
		if graph[makePoint(x, y-1)].value == graph[vertex].value {
			n = append(n, makePoint(x, y-1))
		} else {
			p += 1
		}
	} else {
		p += 1
	}
	if (y+1) > -1 && (y+1) < maxX {
		if graph[makePoint(x, y+1)].value == graph[vertex].value {
			n = append(n, makePoint(x, y+1))
		} else {
			p += 1
		}
	} else {
		p += 1
	}
	return n, p
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
