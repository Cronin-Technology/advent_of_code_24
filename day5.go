package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type page struct {
	x int
	y int
	p bool //used or not
	c int  //if used how many times
}

type print struct {
	value []int
	mid   int
	valid bool
	v     []page // Items to repair
}

var pages = make(map[string]page)
var prints = make(map[int]print)
var rejects = make(map[int]print)
var count int
var count2 int
var count3 int

func main() {
	readFile, err := os.Open("day5_1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		a, b := getPage(fileScanner.Text())
		p := page{a, b, false, 0}
		pages[fileScanner.Text()] = p
	}
	readFile2, err := os.Open("day5_2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	x := 0
	for fileScanner2.Scan() {
		s := strings.Split(fileScanner2.Text(), ",")
		var p []int
		for i := range len(s) {
			a, _ := strconv.Atoi(s[i])
			p = append(p, a)
		}
		pr := print{p, 0, true, []page{}}
		prints[x] = pr
		x++
	}
	checkPrint()
	fixPrint()
	midPoint()
	// fmt.Println(pages)
	// fmt.Println("+++++++++++++++++++++")
	//fmt.Println(prints)
	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", count3)
}

func checkPrint() {
	var toRemove []int
	var toViolate []page
	for y := range len(prints) {
		z := prints[y]
		for k, v := range pages {
			i := slices.Index(z.value, v.x)
			j := slices.Index(z.value, v.y)
			if i != (-1) && j != (-1) {
				vc := v.c + 1
				p := page{v.x, v.y, true, vc}
				pages[k] = p
				if i < j && z.valid {

				} else {
					toRemove = append(toRemove, y)
					toViolate = append(toViolate, p)
				}
			}
		}
		for n := range len(toRemove) {
			d := prints[toRemove[n]]
			newPrint := print{d.value, d.mid, false, toViolate}
			updatePrint(toRemove[n], newPrint)
		}
		toRemove = []int{}
		toViolate = []page{}
	}

}

func checkOnePrint(z print) bool {
	for _, v := range pages {
		i := slices.Index(z.value, v.x)
		j := slices.Index(z.value, v.y)
		if i != (-1) && j != (-1) {
			if i < j {

			} else {
				return false
			}
		}
	}
	return true
}

func fixPrint() {
	for y := range len(prints) {
		v := rejects[y]
		fmt.Println(v)
		temp := v.value
		var fixer []page
		for u := range len(v.v) {
			fixer = append(fixer, v.v[u])
		}
		// sort.Slice(fixer, func(i, j int) bool {
		// 	return fixer[i].y > fixer[j].y
		// })
		// sort.Slice(fixer, func(i, j int) bool {
		// 	return fixer[i].x < fixer[j].x
		// })
		fmt.Println("before:", fixer)
		for _, i := range pages {
			x := slices.Index(temp, i.x)
			y := slices.Index(temp, i.y)
			if x != -1 && y != -1 {
				if x < y {
					temp[x], temp[y] = temp[y], temp[x]
				}
			}
			//p := print{temp, 0, false, fixer}
			//fmt.Println("Iteration: ", i, " :", fixer[i].x, fixer[i].y, temp)
			//fmt.Println("X Index: ", slices.Index(temp, fixer[i].x), "   Y Index: ", slices.Index(temp, fixer[i].y))
		}
		//fmt.Println("after", fixer)
		middleIndex := len(temp) / 2
		if len(temp)%2 == 0 {
			fmt.Println("No Middle Value...")
		} else {
			count3 += temp[middleIndex]
		}
		//fmt.Println("Examples for Checking Reject: ", temp)
	}
}

func midPoint() {
	for _, v := range prints {
		middleIndex := len(v.value) / 2
		if len(v.value)%2 == 0 {
			fmt.Println("No Middle Value...")
		} else {
			if v.valid {
				count += v.value[middleIndex]
			}
		}
	}
	for _, g := range rejects {
		middleIndex := len(g.value) / 2
		if len(g.value)%2 == 0 {
			fmt.Println("No Middle Value...!!!")
		} else {
			if !g.valid {
				count2 += g.value[middleIndex]
			}
		}
	}
}

func updatePrint(i int, p print) {
	prints[i] = p
	rejects[i] = p
}

func getPage(a string) (int, int) {
	b := strings.Split(a, "|")
	bx, _ := strconv.Atoi(b[0])
	by, _ := strconv.Atoi(b[1])
	return bx, by
}

func makePage(x int, y int) string {
	return (strconv.Itoa(x) + "|" + strconv.Itoa(y))
}
