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
	sum2 := 0
	for fileScanner.Scan() {
		s := fileScanner.Text()
		input := Part_Two(s)
		r := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)")
		a := r.FindAllString(s, -1)
		for i := range len(a) {
			sum += Junk(a[i])
		}
		b := r.FindAllString(input, -1)
		for i := range len(b) {
			sum2 += Junk(b[i])
		}
	}
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}

func Junk(s string) int {
	s1 := strings.ReplaceAll(s, "mul(", "")
	s2 := strings.ReplaceAll(s1, ")", "")
	s3 := strings.Split(s2, ",")
	a, _ := strconv.Atoi(s3[0])
	b, _ := strconv.Atoi(s3[1])
	return a * b
}

func Part_Two(input string) string {
	input = "do()" + input
	rdo := regexp.MustCompile("(?i)do\\(\\)")
	rdont := regexp.MustCompile("(?i)don't\\(\\)")
	dos := rdo.FindAllStringSubmatchIndex(input, -1)
	var d []int
	donts := rdont.FindAllStringSubmatchIndex(input, -1)
	var e []int
	for a := range len(dos) {
		d = append(d, dos[a][1])
	}
	for b := range len(donts) {
		e = append(e, donts[b][0])
	}
	fmt.Println(d, e)
	//re := regexp.MustCompile("(?i)mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	//fmt.Println(re.FindAllString(input, -1))
	var s []string
	prev := 0
	for i := 0; i < len(d); i++ {
		j := i
		z := 0
		for z != 1 {
			if i >= len(d) || j >= len(e) {
				if d[len(d)-1] > e[len(e)-1] {
					s = append(s, input[d[i]:len(input)])
				}
				i += len(d)
				z = 1
			} else if d[i] <= prev { //checks if there are duplicate values
				i += 1
			} else if e[j] <= d[i] { //checks if there are multiple do in a row
				j += 1
			} else if e[j] == d[i] { //checks if there are multiple donts in a row
				j += 1
				i += 1
			} else {
				fmt.Println(d[i], e[j])
				s = append(s, input[d[i]:e[j]])
				prev = e[j]
				z = 1
			}
		}
		z = 0
	}

	output := ""
	for x := range len(s) {
		output += s[x]
	}
	fmt.Println("Output:", output)
	return output
}
