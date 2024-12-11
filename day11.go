package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		s := strings.Split(fileScanner.Text(), " ")
		for i := range len(s) {
			a = append(a, s[i])
		}
	}

	fmt.Println(a)
}
