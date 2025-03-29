package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	readFile, err := os.Open("input/day12.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	dir := []string{}
	mov := []int{}
	for fileScanner.Scan() {
		a := fileScanner.Text()
		dir = append(dir, string(a[0]))
		b, _ := strconv.Atoi(a[1:len(a)])
		mov = append(mov, b)
	}
	//fmt.Println(dir, mov)
	fmt.Println("Day 12 Part One Answer: ", dayTwelve_PartOne(dir, mov))
	fmt.Println("Day 12 Part Two Answer: ", dayTwelve_PartTwo(dir, mov))
}

func dayTwelve_PartOne(dir []string, mov []int) int {
	result := 0
	directions := []int{0, 90, 180, 270}
	opposite := []int{180, 270, 0, 90}
	movement := []int{0, 0, 0, 0}
	cardinal := []string{"N", "E", "S", "W"}
	opc := []string{"S", "W", "N", "E"}
	cd := 90
	for i := range dir {
		if dir[i] == "F" {
			a := slices.Index(directions, cd)
			b := slices.Index(opposite, cd)
			if movement[b]-mov[i] <= 0 {
				movement[a] = movement[a] + (-1 * (movement[b] - mov[i]))
				movement[b] = 0
			} else {
				movement[b] = movement[b] - mov[i]
				movement[a] = 0
			}
			//fmt.Println("Forward", dir[i], a, b, cd, mov[i], "Placement: ", movement)
		} else if dir[i] == "L" {
			//pre := cd
			cd = cd - mov[i]
			if cd < 0 {
				cd = cd + 360
			}
			//fmt.Println("Turn Left", pre, mov[i], cd, "Placement: ", movement)
		} else if dir[i] == "R" {
			//pre := cd
			cd = cd + mov[i]
			if cd > 270 {
				cd = cd - 360
			}
			//fmt.Println("Turn Right", pre, mov[i], cd, "Placement: ", movement)
		} else {
			//movement[slices.Index(cardinal, dir[i])] += mov[i]
			a := slices.Index(cardinal, dir[i])
			b := slices.Index(opc, dir[i])
			if movement[b]-mov[i] <= 0 {
				movement[a] = movement[a] + (-1 * (movement[b] - mov[i]))
				movement[b] = 0
			} else {
				movement[b] = movement[b] - mov[i]
				movement[a] = 0
			}
			//fmt.Println("Move", dir[i], a, b, mov[i], "Placement: ", movement)
		}
		//fmt.Println(movement)
	}
	result = sumSlice(movement)
	return result
}

func sumSlice(a []int) int {
	result := 0
	for i := range a {
		result += a[i]
	}
	return result
}

func dayTwelve_PartTwo(dir []string, mov []int) int {
	result := 0
	directions := []int{0, 90, 180, 270}
	opposite := []int{180, 270, 0, 90}
	movement := []int{0, 0, 0, 0}
	waypoint := []int{1, 10, 0, 0}
	cardinal := []string{"N", "E", "S", "W"}
	opc := []string{"S", "W", "N", "E"}
	cd := 90
	for i := range dir {
		if dir[i] == "F" {
			full := []int{}
			for x := range waypoint {
				full = append(full, waypoint[x]*mov[i])
			}
			for y := 0; y < len(full)-1; y++ {
				if full[y] != 0 {
					a := y
					c := directions[y]
					b := slices.Index(opposite, c)
					fmt.Println(a, b)
					if movement[b]-full[y] <= 0 {
						movement[a] = movement[a] + (-1 * (movement[b] - full[y]))
						movement[b] = 0
					} else {
						movement[b] = movement[b] - full[y]
						movement[a] = 0
					}
				}
			}
			fmt.Println("Forward", dir[i], mov[i], "Placement: ", movement, "Waypoint:", waypoint)
		} else if dir[i] == "L" {
			pre := cd
			cd = cd - mov[i]
			if cd < 0 {
				cd = cd + 360
			}
			fmt.Println("Turn Left", pre, mov[i], cd, "Placement: ", movement)
		} else if dir[i] == "R" {

			fmt.Println("Turn Right", "Placement After Rotation: ", waypoint)
		} else {
			a := slices.Index(cardinal, dir[i])
			b := slices.Index(opc, dir[i])
			if waypoint[b]-mov[i] <= 0 {
				waypoint[a] = waypoint[a] + (-1 * (waypoint[b] - mov[i]))
				waypoint[b] = 0
			} else {
				waypoint[b] = waypoint[b] - mov[i]
				waypoint[a] = 0
			}
			fmt.Println("Move", dir[i], a, b, mov[i], "Waypoint: ", waypoint)
		}
		fmt.Println(movement)
	}
	result = sumSlice(movement)
	return result
}
