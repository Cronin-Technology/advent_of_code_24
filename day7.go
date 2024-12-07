package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	result := 0
	result2 := 0
	a := make(map[int][]int)
	readFile, err := os.Open("day7.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), ":")
		v, _ := strconv.Atoi(s[0])
		z := []int{}
		s1 := strings.Split(s[1], " ")
		for i := 1; i < len(s1); i++ {
			c, _ := strconv.Atoi(s1[i])
			z = append(z, c)
		}
		a[v] = z
		result += elephantMath(v, z)
		result2 += elephantMathString(v, z)
	}

	fmt.Println("Part 1: ", result)
	fmt.Println("Part 2: ", result2)
}

func addThem(a int, b []int) bool {
	s := 0
	for i := range len(b) {
		s += b[i]
	}
	if a == s {
		return true
	}
	return false
}

func multiplyThem(a int, b []int) bool {
	s := 0
	for i := range len(b) - 1 {
		s += b[i] * b[i+1]
	}
	if a == s {
		return true
	}
	return false
}

func findFactors(num int) []int {
	a := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			a = append(a, i)
		}
	}
	return a
}

func elephantMath(a int, b []int) int {
	s := b[0] + b[1]
	t := b[0] * b[1]
	root := []int{s, t}
	current := root
	update := []int{}
	result := []int{}
	result = append(result, s)
	result = append(result, t)

	loop := 2

	for loop < len(b) {
		for i := range len(current) {
			if loop < len(b) {
				c := current[i] + b[loop]
				d := current[i] * b[loop]
				update = append(update, c)
				update = append(update, d)
			}
		}
		for j := range len(update) {
			result = append(result, update[j])
		}
		current = update
		loop += 1
	}
	r := 0
	if len(b) >= 2 {
		r = int(math.Pow(float64(2), float64(len(b)-1)))
	} else {
		r = 0
	}
	re := result[len(result)-r : len(result)]

	if slices.Contains(re, a) {
		return a
	}
	return 0
}

func elephantMathString(a int, b []int) int {
	s := b[0] + b[1]
	t := b[0] * b[1]
	t2, _ := strconv.Atoi(strconv.Itoa(b[0]) + strconv.Itoa(b[1]))
	root := []int{s, t, t2}
	current := root
	update := []int{}
	result := []int{}
	result = append(result, s)
	result = append(result, t)
	result = append(result, t2)
	loop := 2

	for loop < len(b) {
		for i := range len(current) {
			if loop < len(b) {
				c := current[i] + b[loop]
				d := current[i] * b[loop]
				e, _ := strconv.Atoi(strconv.Itoa(current[i]) + strconv.Itoa(b[loop]))
				update = append(update, c)
				update = append(update, d)
				update = append(update, e)
			}
		}
		for j := range len(update) {
			result = append(result, update[j])
		}
		current = update
		loop += 1
	}
	r := 0
	if len(b) >= 2 {
		r = int(math.Pow(float64(3), float64(len(b)-1)))
	} else {
		r = 0
	}
	re := result[len(result)-r : len(result)]

	if slices.Contains(re, a) {
		return a
	}
	return 0
}

func sumFactors(a int, f []int, b []int) {
	for i := range len(b) {
		s1 := sumWithForLoop(b[0:i])
		s2 := sumWithForLoop(b[i:len(b)])
		//fmt.Println(s1, s2)
		if slices.Contains(f, s1) && slices.Contains(f, s2) {
			fmt.Println(a)
		}
		if s1*s2 == a {
			fmt.Println(a)
		}
	}
}

func sumWithForLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
