package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type loci struct {
	value string
	x     int
	y     int
	p     bool
	d     []string
	c     int
}

var maxX int
var maxY int

var location = make(map[string]loci)
var l = make(map[string]loci)

func main() {

	part1 := 0
	part2 := 0
	readFile, err := os.Open("day8.txt")
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
			b := loci{s[i], x, y, false, []string{}, 0}
			location[a] = b
			l[a] = b
			x += 1
		}
		maxX = x
		y += 1
	}
	maxY = y

	readFile.Close()
	getAntinodes()
	getPairs()

	for i := range maxX {
		con := ""
		for j := range maxY {
			a := makePoint(j, i)
			if location[a].p {
				con += "#"
				part1 += 1
			} else {
				con += location[a].value
			}
		}
		fmt.Println(con)
	}

	fmt.Println("")
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("")

	for i := range maxX {
		con := ""
		for j := range maxY {
			a := makePoint(j, i)
			if l[a].p {
				con += "#"
				part2 += 1
			} else {
				con += location[a].value
			}
		}
		fmt.Println(con)
	}

	fmt.Println("Part 1: ", part1)

	fmt.Println("Part 2: ", part2)
}

func getPairs() {
	a := make(map[string][]string)
	temp := []string{}
	unique := []string{}
	//pair := [][2]string{}
	for _, v := range l {
		if !slices.Contains(unique, v.value) && v.value != "." {
			unique = append(unique, v.value)
		}
	}
	//fmt.Println(unique)
	for i := range len(unique) {
		for k, v := range l {
			if v.value == unique[i] {
				temp = append(temp, k)
			}
		}
		a[unique[i]] = temp
		temp = []string{}
	}

	for key, value := range a {
		p := makePairs(value)
		for i := range len(p) {
			slope := []int{}
			r := 0
			q := 1
			x1, y1 := getPoint(p[i][0])
			x2, y2 := getPoint(p[i][1])
			// if y1 > y2 {
			// 	x1, y1 = getPoint(p[i][1])
			// 	x2, y2 = getPoint(p[i][0])
			// }
			slope = append(slope, x2-x1)
			slope = append(slope, y2-y1)
			fmt.Println(key, x1, y1, x2, y2, "::", slope)
			for ok := true; ok; ok = (r != 2) {
				f := makePoint(x1-(q*slope[0]), y1-(q*slope[0]))
				if x1-(q*slope[0]) >= 0 && x1-(q*slope[0]) <= maxX && y1-(q*slope[0]) >= 0 && y1-(q*slope[0]) <= maxY {
					l[f] = loci{"#", x1 - (q * slope[0]), y1 - (q * slope[0]), true, []string{}, 1}
				} else {
					r = 2
				}
				q += 1
			}
			r = 0
			q = 1
			for ok := true; ok; ok = (r != 2) {
				f := makePoint(x2+(q*slope[1]), y2+(q*slope[1]))
				if x2+(q*slope[1]) >= 0 && x2+(q*slope[1]) <= maxX && y2+(q*slope[1]) >= 0 && y2+(q*slope[1]) <= maxY {
					l[f] = loci{"#", x2 + (q * slope[1]), y2 + (q * slope[1]), true, []string{}, 1}
				} else {
					r = 2
				}
				q += 1
			}
		}
	}
}

func getAntinodes() {
	a := make(map[string][]string)
	temp := []string{}
	unique := []string{}
	//pair := [][2]string{}
	for _, v := range location {
		if !slices.Contains(unique, v.value) && v.value != "." {
			unique = append(unique, v.value)
		}
	}
	//fmt.Println(unique)
	for i := range len(unique) {
		for k, v := range location {
			if v.value == unique[i] {
				temp = append(temp, k)
			}
		}
		a[unique[i]] = temp
		temp = []string{}
	}

	for _, value := range a {
		// fmt.Println(key, ":", value)
		//fmt.Println(key, ":", makePairs(value))
		p := makePairs(value)
		for i := range len(p) {
			x1, y1 := getPoint(p[i][0])
			x2, y2 := getPoint(p[i][1])
			xt := getAbsInt(x1, x2) * 2
			yt := getAbsInt(y1, y2) * 2
			//fmt.Println(x1, y1, x2, y2, ":", xt, yt)
			if y1 < y2 {
				if x1 > x2 {
					w := location[makePoint(x1-xt, y1+yt)]
					q := location[makePoint(x2+xt, y2-yt)]
					location[makePoint(x1-xt, y1+yt)] = loci{w.value, w.x, w.y, true, w.d, w.c}
					location[makePoint(x2+xt, y2-yt)] = loci{q.value, q.x, q.y, true, q.d, q.c}
				} else { //x1 < x2
					w := location[makePoint(x1+xt, y1+yt)]
					q := location[makePoint(x2-xt, y2-yt)]
					location[makePoint(x1+xt, y1+yt)] = loci{w.value, w.x, w.y, true, w.d, w.c}
					location[makePoint(x2-xt, y2-yt)] = loci{q.value, q.x, q.y, true, q.d, q.c}
				}
			} else if y1 > y2 {
				if x1 > x2 {
					w := location[makePoint(x1-xt, y1-yt)]
					q := location[makePoint(x2+xt, y2+yt)]
					location[makePoint(x1-xt, y1-yt)] = loci{w.value, w.x, w.y, true, w.d, w.c}
					location[makePoint(x2+xt, y2+yt)] = loci{q.value, q.x, q.y, true, q.d, q.c}
				} else { //x1 < x2
					w := location[makePoint(x1+xt, y1-yt)]
					q := location[makePoint(x2-xt, y2+yt)]
					location[makePoint(x1+xt, y1-yt)] = loci{w.value, w.x, w.y, true, w.d, w.c}
					location[makePoint(x2-xt, y2+yt)] = loci{q.value, q.x, q.y, true, q.d, q.c}
				}
			}
		}
	}

}

func getAbsInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func makePairs(strings []string) [][2]string {
	var pairs [][2]string

	// Iterate over the slice
	for i := 0; i < len(strings); i++ {
		for j := i + 1; j < len(strings); j++ {
			// Add the pair (strings[i], strings[j])
			pairs = append(pairs, [2]string{strings[i], strings[j]})
		}
	}

	return pairs
}

// func getPoint(a string) (int, int) {
// 	b := strings.Split(a, ":")
// 	bx, _ := strconv.Atoi(b[0])
// 	by, _ := strconv.Atoi(b[1])
// 	return bx, by
// }

// func makePoint(x int, y int) string {
// 	return (strconv.Itoa(x) + ":" + strconv.Itoa(y))
// }
