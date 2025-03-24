package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
	cnt                                    int
	valid                                  bool
}

func main() {
	readFile, err := os.Open("input/day4.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	validity := []Passport{}
	var lines []string
	cline := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			lines = append(lines, cline)
			cline = ""
		}
		cline += (" " + line)
	}
	for l := range lines {
		p := Passport{}
		cnt := 0
		a := strings.Split(lines[l], " ")
		for i := range a {
			b := strings.Split(a[i], ":")
			if b[0] == "byr" {
				p.byr = b[1]
				cnt++
			}
			if b[0] == "iyr" {
				p.iyr = b[1]
				cnt++
			}
			if b[0] == "eyr" {
				p.eyr = b[1]
				cnt++
			}
			if b[0] == "hgt" {
				p.hgt = b[1]
				cnt++
			}
			if b[0] == "hcl" {
				p.hcl = b[1]
				cnt++
			}
			if b[0] == "ecl" {
				p.ecl = b[1]
				cnt++
			}
			if b[0] == "pid" {
				p.pid = b[1]
				cnt++
			}
			if b[0] == "cid" {
				p.cid = b[1]
				cnt++
			}
		}
		p.cnt = cnt
		validity = append(validity, p)
	}
	// for u := range lines {
	// 	fmt.Println(lines[u])
	// }
	fmt.Println("Day Four Part One Answer: ", dayFour_PartOne(validity))
	fmt.Println("Day Four Part Two Answer: ", dayFour_PartTwo(validity))
}

func dayFour_PartOne(validity []Passport) int {
	result := 0
	for x := range validity {
		if validity[x].cnt == 8 {
			result++
			validity[x].valid = true
		} else if validity[x].cnt == 7 && validity[x].cid == "" {
			result++
			validity[x].valid = true
		} else {
			validity[x].valid = false
		}
	}
	return result
}

func dayFour_PartTwo(validity []Passport) int {
	result := 0
	for x := range validity {
		if validity[x].cnt == 8 {
			validity[x].valid = validatePassport(validity[x])
			if validity[x].valid {
				result++
			}
		} else if validity[x].cnt == 7 && validity[x].cid == "" {
			validity[x].valid = validatePassport(validity[x])
			if validity[x].valid {
				result++
			}
		} else {
			validity[x].valid = false
		}
	}
	return result
}

func validatePassport(p Passport) bool {
	eyecolor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	haircolor := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	byr, _ := strconv.Atoi(p.byr)
	byrRange := []int{1920, 2002}
	iyr, _ := strconv.Atoi(p.iyr)
	iyrRange := []int{2010, 2020}
	eyr, _ := strconv.Atoi(p.eyr)
	eyrRange := []int{2020, 2030}
	heights := []string{"cm", "in"}
	hgtRangeInch := []int{59, 76}
	hgtRangeCm := []int{150, 193}
	result := 0
	if byr >= byrRange[0] && byr <= byrRange[1] {
		result++
	} else {
		//fmt.Println("byr Fail: ", p.byr)
	}
	if iyr >= iyrRange[0] && iyr <= iyrRange[1] {
		result++
	} else {
		//fmt.Println("iyr Fail: ", p.iyr)
	}
	if eyr >= eyrRange[0] && eyr <= eyrRange[1] {
		result++
	} else {
		//fmt.Println("Eyr Fail: ", p.eyr)
	}
	//fmt.Println("byr:", byr, " iyr:", iyr, " eyr:", eyr)
	if strings.Contains(p.hgt, heights[0]) { //cm
		hs := strings.Split(p.hgt, heights[0])
		hsn, _ := strconv.Atoi(hs[0])
		if hsn >= hgtRangeCm[0] && hsn <= hgtRangeCm[1] {
			result++
			//fmt.Println("Height in CM: ", p.hgt)
		} else {
			//fmt.Println("Height Fail: ", p.hgt)
		}
	} else if strings.Contains(p.hgt, heights[1]) { //inch
		hs := strings.Split(p.hgt, heights[1])
		hsn, _ := strconv.Atoi(hs[0])
		if hsn >= hgtRangeInch[0] && hsn <= hgtRangeInch[1] {
			result++
			//fmt.Println("Height in Inch: ", p.hgt)
		} else {
			//fmt.Println("Height Fail: ", p.hgt)
		}
	}
	if string(p.hcl[0]) == "#" {
		if len(p.hcl) == 7 {
			for i := 1; i < len(p.hcl); i++ {
				if slices.Contains(haircolor, string(p.hcl[i])) {
					//do nothing
				} else {
					//fmt.Println("Haircolor Fail: ", p.hcl)
					break
				}
			}
			result++
		}
	}
	if slices.Contains(eyecolor, p.ecl) {
		result++
	} else {
		//fmt.Println("Eyecolor Fail:", p.ecl)
	}
	if len(p.pid) == 9 {
		_, err := strconv.Atoi(p.pid)
		if err != nil {
			//fmt.Println("PID Fail: ", p.pid)
		} else {
			result++
		}
	}

	if result == 7 {
		return true
	} else {
		//fmt.Println(p, " ENTRY FAILED ", result)
		return false
	}
}
