package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var a [][]int
	readFile, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var b []int
		//a = append(a, fileScanner.Text())
		s := strings.Split(fileScanner.Text(), " ")
		for i := range len(s) {
			c, _ := strconv.Atoi(s[i])
			b = append(b, c)
		}
		a = append(a, b)
	}

	//CheckReports(a)
	// for f := range len(a) {
	// 	fmt.Println(a[f])
	// }

	partone := 0
	parttwo := 0

	for i := range a {
		if CheckData(a[i]) {
			partone += 1
		}
	}

	for i := range a {
		if CheckDataTwice(a[i]) {
			parttwo += 1
		}
	}

	fmt.Println("Part 1: ", partone)
	fmt.Println("Part 2: ", partone+parttwo)
}

func CheckDataAgain(data []int, index int) bool {
	d := append([]int(nil), data...)
	e := append([]int(nil), data...)
	f := append([]int(nil), data...)
	if index > 0 && index < len(data) {
		d = slices.Delete(d, index, index+1)
		e = slices.Delete(e, index+1, index+2)
		f = slices.Delete(f, index-1, index)
	} else if index == 0 {
		d = slices.Delete(d, index, index+1)
		e = slices.Delete(e, index+1, index+2)
		f = slices.Delete(f, index, index+1)
	} else if index == len(data)-1 {
		d = slices.Delete(d, index-1, index)
		e = slices.Delete(e, index-2, index)
		f = slices.Delete(f, index-1, index)
	}

	if CheckData(d) || CheckData(e) || CheckData(f) {
		return true
	} else {
		fmt.Println(data, d, e, f, index)
		//fmt.Println(data, ":", d, ":", e, " ", index)
		return false
	}
}

func CheckDataTwice(data []int) bool {
	count := 0
	if data[0]-data[1] < 0 {
		for j := 0; j < len(data)-1; j++ {
			if data[j]-data[j+1] < 0 && data[j]-data[j+1] > -4 && data[j]-data[j+1] != 0 {
				count += 0
			} else if CheckDataAgain(data, j) {
				return true
			} else {
				return false
			}
		}
	} else if data[0]-data[1] > 0 {
		for j := 0; j < len(data)-1; j++ {
			if data[j]-data[j+1] > 0 && data[j]-data[j+1] < 4 && data[j]-data[j+1] != 0 {
				count += 0
			} else if CheckDataAgain(data, j) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func CheckData(data []int) bool {
	count := 0
	if data[0]-data[1] < 0 {
		for j := 0; j < len(data)-1; j++ {
			if data[j]-data[j+1] < 0 && data[j]-data[j+1] > -4 && data[j]-data[j+1] != 0 {
				count += 0
			} else {
				return false
			}
		}
		return true
	} else if data[0]-data[1] > 0 {
		for j := 0; j < len(data)-1; j++ {
			if data[j]-data[j+1] > 0 && data[j]-data[j+1] < 4 && data[j]-data[j+1] != 0 {
				count += 0
			} else {
				return false
			}
		}
		return true
	}
	return false
}
