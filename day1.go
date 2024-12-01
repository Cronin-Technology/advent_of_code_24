package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var a []int
	var b []int
	c := 0
	d := 0
	readFile, err := os.Open("day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), " ")
		aInt, _ := strconv.Atoi(s[0])
		bInt, _ := strconv.Atoi(s[3])
		a = append(a, aInt)
		b = append(b, bInt)
	}

	readFile.Close()

	sort.Ints(a)
	sort.Ints(b)

	for i := range len(a) {
		if (a[i] - b[i]) < 0 {
			c = c + (-1 * (a[i] - b[i]))
		} else if (a[i] - b[i]) > 0 {
			c = c + (a[i] - b[i])
		} else {
			c = c + 0
		}

	}

	for j := range len(a) {
		valueToCount := a[j]
		count := 0

		for _, num := range b {
			if num == valueToCount {
				count++
			}
		}

		d = d + (a[j] * count)

	}

	fmt.Println("Part 1: ", c)
	fmt.Println("Part 2: ", d)
}
