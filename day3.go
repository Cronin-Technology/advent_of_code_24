package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("day3.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	sum := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		r := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)")
		a := r.FindAllString(s, 10000000)
		for i := range len(a) {
			sum += Junk(a[i])
		}

	}
	fmt.Println("Part 1: ", sum)
}

func Junk(s string) int {
	s1 := strings.ReplaceAll(s, "mul(", "")
	s2 := strings.ReplaceAll(s1, ")", "")
	s3 := strings.Split(s2, ",")
	a, _ := strconv.Atoi(s3[0])
	b, _ := strconv.Atoi(s3[1])
	return a * b
}
