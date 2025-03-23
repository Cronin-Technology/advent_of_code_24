package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	value              string
	min, max, cnt, pos int
	pass               string
}

func main() {
	readFile, err := os.Open("input/day2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	validity := []Policy{}
	for fileScanner.Scan() {
		p := Policy{}
		a := strings.Split(fileScanner.Text(), ":")
		b := strings.Split(a[0], " ")
		c := strings.Split(b[0], "-")
		p.min, _ = strconv.Atoi(c[0])
		p.max, _ = strconv.Atoi(c[1])
		p.value = b[1]
		p.pass = a[1]
		validity = append(validity, p)
	}
	fmt.Println("Day Two Part One Answer: ", dayTwo_PartOne(validity))
	fmt.Println("Day Two Part Two Answer: ", dayTwo_PartTwo(validity))
}

func dayTwo_PartOne(validity []Policy) int {
	result := 0
	for x := range validity {
		validity[x].cnt = strings.Count(validity[x].pass, validity[x].value)
		if validity[x].cnt >= validity[x].min && validity[x].cnt <= validity[x].max {
			result++
		}
	}
	return result
}

func dayTwo_PartTwo(validity []Policy) int {
	result := 0
	for x := range validity {
		tmp := 0
		if string(validity[x].pass[validity[x].min]) == validity[x].value {
			tmp++
		}
		if string(validity[x].pass[validity[x].max]) == validity[x].value {
			tmp++
		}
		if tmp == 1 {
			result++
		}
	}
	return result
}
