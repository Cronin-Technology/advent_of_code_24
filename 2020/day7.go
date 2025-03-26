package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input/day7.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	fmt.Println("Step 1:")
	contains, containedBy := bagMaps(lines)
	fmt.Println(step1(containedBy))

	fmt.Println("Step 2:")
	fmt.Println(step2(contains))
}

type bagCount struct {
	color string
	num   int
}

var (
	inputRegex = regexp.MustCompile("^([a-z ]+) bags contain ([a-z0-9, ]+)\\.$")
	bagRegex   = regexp.MustCompile("^(\\d+) ([a-z ]+) bag[s]?$")
)

func bagMaps(lines []string) (map[string][]bagCount, map[string][]string) {
	contains, containedBy := make(map[string][]bagCount), make(map[string][]string)
	for _, l := range lines {
		tokens := inputRegex.FindStringSubmatch(l)
		if tokens == nil {
			log.Fatalf("Failed to parse %q\n", l)
		}
		container := tokens[1]
		for _, c := range strings.Split(tokens[2], ", ") {
			if c == "no other bags" {
				continue
			}
			contents := bagRegex.FindStringSubmatch(c)
			if contents == nil {
				log.Fatalf("Failed to parse %q\n", c)
			}
			bag := bagCount{}
			qty, _ := strconv.Atoi(contents[1])
			bag.num = qty
			bag.color = contents[2]
			contains[container] = append(contains[container], bag)
			containedBy[bag.color] = append(containedBy[bag.color], container)
		}
	}
	return contains, containedBy
}

const targetBag = "shiny gold"

func step1(containedBy map[string][]string) int {
	canContain := containedBy[targetBag]
	seen := make(map[string]bool)
	seen[targetBag] = true
	for len(canContain) > 0 {
		curr := canContain[0]
		canContain = canContain[1:]
		if seen[curr] {
			continue
		}
		seen[curr] = true
		canContain = append(canContain, containedBy[curr]...)
	}
	// Subtract one for shiny gold
	return len(seen) - 1
}

func step2(contains map[string][]bagCount) int64 {
	return countContents(targetBag, contains)
}

func countContents(target string, contains map[string][]bagCount) int64 {
	sum := int64(0)
	for _, c := range contains[target] {
		sum += int64(c.num)
		sum += int64(c.num) * countContents(c.color, contains)
	}
	return sum
}
