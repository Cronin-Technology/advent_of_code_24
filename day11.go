package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := []string{}
	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		s := strings.Split(fileScanner.Text(), " ")
		for i := range len(s) {
			a = append(a, s[i])
		}
	}
	now := a
	nows := [][]string{}
	for i := 0; i < len(now); i++ {
		insert := []string{now[i]}
		nows = append(nows, insert)
	}

	for i := 0; i < 75; i++ {
		//fmt.Println(len(now))
		fmt.Println("Iteration", i)
		now = Blink(now)
	}
	count := 0
	// for i := 0; i < len(nows); i++ {
	// 	//fmt.Println(now)
	// 	for j := 0; j < 40; j++ {
	// 		nows[i] = Blink(nows[i])
	// 	}
	// 	fmt.Println("Combined Iteration", i)
	// }

	fmt.Println("Part 1:", len(now))
	fmt.Println("Part 2:", count)
}

func insert(slice []string, index int, value string) []string {
	// Check for valid index
	if index < 0 || index > len(slice) {
		fmt.Println("Index out of bounds")
		return slice
	}
	// Insert the value at the specified index
	slice = append(slice[:index], append([]string{value}, slice[index:]...)...)
	return slice
}

func deepCopy(slice []string) []string {
	// Create a new slice with the same length as the original slice
	newSlice := make([]string, len(slice))
	// Copy each element from the original slice to the new slice
	copy(newSlice, slice)
	return newSlice
}

func EvenSlice(t string) (x, y string) {
	ts := strings.Split(t, "")
	t1 := ts[0 : len(t)/2]
	t2 := ts[len(t)/2 : len(t)]
	ta := ""
	tb := ""
	for i := range len(t1) {
		ta += t1[i]
		tb += t2[i]
	}
	tx, _ := strconv.Atoi(ta)
	ty, _ := strconv.Atoi(tb)
	if tx == 0 {
		tx = 0
	}
	ta = strconv.Itoa(tx)
	tb = strconv.Itoa(ty)
	return ta, tb
}

func Blink(a []string) []string {
	temp := []string{}
	for i := range len(a) {
		if a[i] == "0" {
			temp = append(temp, "1")
		} else if len(a[i])%2 == 0 {
			x, y := EvenSlice(a[i])
			temp = append(temp, x)
			temp = append(temp, y)
		} else {
			z, _ := strconv.Atoi(a[i])
			z = z * 2024
			temp = append(temp, strconv.Itoa(z))
		}
	}
	return temp
}

func Blinking(a []string) []string {
	temp := []string{}
	for i := range len(a) {
		if a[i] == "0" {
			temp = append(temp, "1")
		} else if len(a[i])%2 == 0 {
			x, y := EvenSlice(a[i])
			temp = append(temp, x)
			temp = append(temp, y)
		} else {
			z, _ := strconv.Atoi(a[i])
			z = z * 2024
			temp = append(temp, strconv.Itoa(z))
		}
	}
	return temp
}
