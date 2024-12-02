package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a []string
	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		a = append(a, fileScanner.Text())
	}

	UniverseExpansion(a)
	// for i := range len(a) {
	// 	fmt.Println(a[i])
	// }
}

func UniverseExpansion(a []string) {
	//Find columns without galaxies
	var col []int
	var row []int
	for i := range len(a[0]) {
		colCheck := true
		for j := range len(a) {
			if string(a[j][i]) == "#" {
				colCheck = false
			}
		}
		if colCheck {
			col = append(col, i)
		}
	}
	for i := range len(a[0]) {
		rowCheck := true
		for j := range len(a) {
			if string(a[i][j]) == "#" {
				rowCheck = false
			}
		}
		if rowCheck {
			row = append(row, i)
		}
	}
	fmt.Println(col)
	fmt.Println(row)

}
