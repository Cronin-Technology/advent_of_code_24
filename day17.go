package main

import (
	"fmt"
	"math"
)

var a, b, c int

func main() {
	prog := []int{0, 1, 5, 4, 3, 0}
	a = 729
	b = 0
	c = 0
	operand := 0

	program := ChunkProgram(prog, 2)
	fmt.Println(program)
	for i := 0; i < len(program)-1; i++ {
		if program[i][1] <= 3 {
			operand = program[i][1]
		} else if program[i][1] == 4 {
			operand = a
		} else if program[i][1] == 5 {
			operand = b
		} else if program[i][1] == 6 {
			operand = c
		} else if program[i][1] == 7 {
			fmt.Println("Failed instruction")
		}

		if program[i][0] == 0 {
			a = int(a / int(math.Pow(2, float64(operand))))
		} else if program[i][0] == 1 {
			b = b ^ program[i][1]
		} else if program[i][0] == 2 {
			b = operand % 8
		} else if program[i][0] == 3 {
			if a != 0 {
				i = operand
			}
		} else if program[i][0] == 4 {
			b = b ^ c
		} else if program[i][0] == 5 {
			fmt.Println(operand % 8)
		} else if program[i][0] == 6 {
			b = int(a / int(math.Pow(2, float64(operand))))
		} else if program[i][0] == 7 {
			c = int(a / int(math.Pow(2, float64(operand))))
		}
		//fmt.Println("Iteration", a, b, c)
	}

}

func ChunkProgram(slice []int, chunkSize int) (chunks [][]int) {
	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}
	return append(chunks, slice)
}
