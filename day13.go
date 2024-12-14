package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type claw struct {
	ax, ay, bx, by, tx, ty int
}

func main() {
	// part2 := 0
	games := []claw{}
	readFile, err := os.Open("day13.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	step := 0
	var ax, ay, bx, by, tx, ty int
	for fileScanner.Scan() {
		s := fileScanner.Text()
		if s == "" {
			games = append(games, claw{ax, ay, bx, by, tx, ty})
			step += 1
		} else {
			if strings.Contains(s, "Button A") {
				ax, ay = parseCoordinates(s)
			}
			if strings.Contains(s, "Button B") {
				bx, by = parseCoordinates(s)
			}
			if strings.Contains(s, "Prize") {
				tx, ty = parsePrizeCoordinates(s)
			}
		}
	}
	sum := 0
	tsum := 0
	for i := range len(games) {
		sum += playGame(games[i])
		//tsum += runGame((games[i]))
	}

	readFile.Close()

	fmt.Println(sum, tsum)
}

func runGame(c claw) int {
	Axy := [][2]int{}
	Bxy := [][2]int{}
	sum := 0
	for i := 0; i < 10000000000000; i++ {
		Axy = append(Axy, [2]int{c.ax * i, c.ay * i})
	}
	for i := 0; i < 10000000000000; i++ {
		Bxy = append(Bxy, [2]int{c.bx * i, c.by * i})
	}
	exit := 0
	s := 0
	for ok := true; ok; ok = (exit != 1) {
		for i := range len(Bxy) {
			for j := range len(Axy) {
				if Bxy[i][0]+Axy[j][0] == c.tx {
					if Bxy[i][1]+Axy[j][1] == c.ty {
						sum += ((i) + j*3)
						s = 1
						exit = 1
						break
					}
				}
			}
		}
		if s == 0 {
			fmt.Println("No Solution Found")
		}
		exit = 1
	}

	return sum
}

func playGame(c claw) int {

	d := (c.ax*c.by - c.ay*c.bx)
	a := (c.tx*c.by - c.ty*c.bx) / d
	b := (c.ax*c.ty - c.ay*c.tx) / d
	if ((c.ax*a)+(c.bx*b)) == c.tx && ((c.ay*a)+(c.by*b)) == c.ty {
		return ((3 * a) + b)
	}
	return 0
}

func parseCoordinates(input string) (int, int) {
	re := regexp.MustCompile(`X([+-]?\d+), Y([+-]?\d+)`)
	matches := re.FindStringSubmatch(input)

	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return x, y
}

func parsePrizeCoordinates(input string) (int, int) {
	re := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	matches := re.FindStringSubmatch(input)

	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	//Part 1 : >> do this
	//return x, y
	return (x + 10000000000000), (y + 10000000000000)
}
