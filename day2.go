package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a [][]int
	readFile, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var b []int
		//a = append(a, fileScanner.Text())
		s := strings.Split(fileScanner.Text(), " ")
		for i := range len(s) {
			c, _ := strconv.Atoi(s[i])
			b = append(b, c)
		}
		a = append(a, b)
	}

	CheckReports(a)
	// for f := range len(a) {
	// 	fmt.Println(a[f])
	// }
}

func CheckReports(a [][]int) {
	total := 0
	modTotal := 0
	for i := range len(a) {
		//direction := ""
		count := 0
		modCount := 0
		if a[i][0]-a[i][1] < 0 {
			//direction = "Increasing"
			for j := range len(a[i]) - 1 {
				if a[i][j]-a[i][j+1] < 0 && a[i][j]-a[i][j+1] > -4 && a[i][j]-a[i][j+1] != 0 {
					count += 0
				} else {
					count += 1
					if count > 0 && j+2 < len(a[i]) && a[i][j]-a[i][j+2] < 0 && a[i][j]-a[i][j+2] > -4 && a[i][j]-a[i][j+2] != 0 {
						modCount += 1
					}
				}
			}
		} else if a[i][0]-a[i][1] > 0 {
			//direction = "Decreasing"
			for j := range len(a[i]) - 1 {
				if a[i][j]-a[i][j+1] > 0 && a[i][j]-a[i][j+1] < 4 && a[i][j]-a[i][j+1] != 0 {
					count += 0
				} else {
					count += 1
					if count > 0 && j+2 < len(a[i]) && a[i][j]-a[i][j+2] > 0 && a[i][j]-a[i][j+2] < 4 && a[i][j]-a[i][j+2] != 0 {
						modCount += 1
					}
				}
			}
		} else if a[i][0]-a[i][1] == 0 {
			count += 1
			if a[i][1]-a[i][2] < 0 {
				//direction = "Increasing"
				for j := range len(a[i]) - 1 {
					if a[i][j]-a[i][j+1] < 0 && a[i][j]-a[i][j+1] > -4 && a[i][j]-a[i][j+1] != 0 {
						count += 0
					} else {
						count += 1
						if count > 0 && j+2 < len(a[i]) && a[i][j]-a[i][j+2] < 0 && a[i][j]-a[i][j+2] > -4 && a[i][j]-a[i][j+2] != 0 {
							modCount += 1
						}
					}
				}
			} else if a[i][1]-a[i][2] > 0 {
				//direction = "Decreasing"
				for j := range len(a[i]) - 1 {
					if a[i][j]-a[i][j+1] > 0 && a[i][j]-a[i][j+1] < 4 && a[i][j]-a[i][j+1] != 0 {
						count += 0
					} else {
						count += 1
						if count > 0 && j+2 < len(a[i]) && a[i][j]-a[i][j+2] > 0 && a[i][j]-a[i][j+2] < 4 && a[i][j]-a[i][j+2] != 0 {
							modCount += 1
						}
					}
				}
			}
		}
		if count == 0 && modCount == 0 {
			total += 1
			//fmt.Println(a[i], " ", count, " direction: ", direction)
		}
		if modCount == 1 {
			modTotal += 1
		}
	}
	fmt.Println("Part 1: ", total)
	fmt.Println("Part 2: ", modTotal, total+modTotal)
}
