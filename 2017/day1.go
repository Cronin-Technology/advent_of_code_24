package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	sum := []int{}
	rsum := []int{}
	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), "")
		s = append(s, s[0])
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				a, _ := strconv.Atoi(s[i])
				sum = append(sum, a)
			}
		}
		s1 := strings.Split(fileScanner.Text(), "")
		for i := 0; i < len(s1); i++ {
			d := (i + len(s1)/2)
			if d > len(s1) {
				d = (i + len(s1)/2) % 2
			}
			if s1[i] == s[d] {
				a, _ := strconv.Atoi(s1[i])
				rsum = append(rsum, a)
			}
		}
	}
	readFile.Close()

	partone := 0
	fmt.Println(rsum)
	for i := range len(rsum) {
		partone += rsum[i]
	}

	fmt.Println("Part 1 : ", partone)
}
