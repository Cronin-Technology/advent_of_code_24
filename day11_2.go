package main

import (
	"fmt"
	"strconv"
	"strings"
)

var mem map[[2]int]int

func rules(num int) []int {
	if num == 0 {
		return []int{1}
	}

	z := strconv.Itoa(num)
	if len(z)%2 == 0 {
		ts := strings.Split(z, "")
		t1 := ts[0 : len(z)/2]
		t2 := ts[len(z)/2 : len(z)]
		ta := ""
		tb := ""
		for i := range len(t1) {
			ta += t1[i]
			tb += t2[i]
		}
		tx, _ := strconv.Atoi(ta)
		ty, _ := strconv.Atoi(tb)
		return []int{tx, ty}
	}

	return []int{num * 2024}

}

func dfs(num, i int) int {
	if i == 0 {
		return 1
	}

	key := [2]int{num, i}
	if val, found := mem[key]; found {
		return val
	}

	count := 0
	arr := rules(num)
	for _, n := range arr {
		count += dfs(n, i-1)
	}

	mem[key] = count
	return count
}

func main() {
	input := []int{1117, 0, 8, 21078, 2389032, 142881, 93, 385} // Example input, replace with your actual input values
	blinks := 75                                                // Example blinks, replace with your actual blinks value

	count := 0
	mem = make(map[[2]int]int)
	for _, num := range input {
		count += dfs(num, blinks)
	}

	fmt.Println(count) // Print the result
}
