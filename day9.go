package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type codon struct {
	i int
	j int
	k int
}

func main() {
	readFile, err := os.Open("day9.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	dna := []codon{}
	rna := []codon{}
	for fileScanner.Scan() {
		input := strings.Split(fileScanner.Text(), "")
		d := ChunkSlice(input, 2)
		for i := 0; i < len(d); i++ {
			if len(d[i]) < 2 {
				a := i
				b, _ := strconv.Atoi(d[i][0])
				c := 0
				dna = append(dna, codon{a, b, c})
				rna = append(rna, codon{a, b, c})
			} else {
				a := i
				b, _ := strconv.Atoi(d[i][0])
				c, _ := strconv.Atoi(d[i][1])
				dna = append(dna, codon{a, b, c})
				rna = append(rna, codon{a, b, c})
			}
		}
	}
	//fmt.Println("DNA:", dna)
	deadspace := GetDeadspace(dna)
	//fmt.Println(deadspace)
	//fmt.Println(ExpressDNA(dna))

	fin := SortDNA(rna, deadspace)
	fmt.Println(ExpressRNA(fin, deadspace))
}

func GetDeadspace(d []codon) int {
	dead := 0
	for i := range len(d) {
		dead += d[i].k
	}
	//fmt.Println(dead)
	return dead
}

func SortDNA(dn []codon, deadspace int) []codon {
	d := dn
	death := 1
	//fmt.Println("START: ", d)
	r := 10
	w := 0
	index := 0
	for ok := true; ok; ok = (r != 0) {
		currentCodon := codon{}
		//get last codon
		lastCodon := d[len(d)-1]
		//get first codon with > 0 in last place
		for ok := true; ok; ok = (w != 1) { //do while loop
			for i := index; i < len(d)-1; i++ {
				if d[i].k != 0 {
					currentCodon = d[i]
					index = i
					w = 1
					break
				}
			}
		}
		//fmt.Println(d)
		//fmt.Println(currentCodon, lastCodon)
		if currentCodon.k-lastCodon.j >= 0 {
			lastCodon.k = currentCodon.k - lastCodon.j
			currentCodon.k = 0
			d[index].k = currentCodon.k
			death += lastCodon.j
			InsertAt(d, index+1, lastCodon)
		} else if currentCodon.k-lastCodon.j < 0 {
			lastCodon.j = lastCodon.j - currentCodon.k
			newCodon := codon{lastCodon.i, currentCodon.k, 0}
			currentCodon.k = 0
			d[index].k = currentCodon.k
			death += newCodon.j
			InsertAt(d, index+1, newCodon)
			d = append(d, lastCodon)
		}
		if death >= deadspace {
			r = 0
			break
		}
	}
	//fmt.Println(d)
	return d
}

func InsertAt(slice []codon, index int, value codon) []codon {
	// Insert the value at the specified index by slicing the slice into two parts
	slice = append(slice[:index], append([]codon{value}, slice[index:]...)...)

	// Drop the last element
	return slice[:len(slice)-1]
}

func ExpressDNA(dna []codon) string {
	exp := ""
	for i := range len(dna) {
		for j := 0; j < dna[i].j; j++ {
			exp += strconv.Itoa(dna[i].i)
		}
		for k := 0; k < dna[i].k; k++ {
			exp += "."
		}
	}
	return exp
}

func ExpressRNA(dna []codon, dead int) (string, int) {
	exp := ""
	s := 0
	for i := range len(dna) {
		for j := 0; j < dna[i].j; j++ {
			exp += strconv.Itoa(dna[i].i)
		}
		for k := 0; k < dna[i].k; k++ {
			exp += "."
		}
	}

	for k := 0; k < len(exp)-1; k++ {
		foo, _ := strconv.Atoi(string(exp[k]))
		s += foo * k
	}

	// for j := 1; j < dead; j++ {
	// 	exp += "."
	// }
	return "", s
}

func ChunkSlice(slice []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}
	return append(chunks, slice)
}
