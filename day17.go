package main

import (
	"fmt"
	"math"
)

var a, b, c int

func main() {
	prog := []int{2, 4, 1, 3, 7, 5, 0, 3, 1, 5, 4, 1, 5, 5, 3, 0}
	found := false
	start := 2024
	output := []int{}
	for !found && start < 636875300000 {
		output = runProgram(start, prog)
		//fmt.Println("Iter:", start)
		if len(output) != len(prog) {
			found = false
			start = start + 1
		} else {
			found = true
		}
		// for x := range output {
		// 	if output[x] != prog[x] {
		// 		//False
		// 		start = start + 2024
		// 		found = false
		// 		break
		// 	}

		// }
	}
	fmt.Println(start, prog, output)

}

func runProgram(input int, prog []int) []int {
	a = input
	b = 0
	c = 0
	i := 0

	out := []int{}

	for i < len(prog) {
		opcode := prog[i]
		operand := prog[i+1]
		combo := getOperand(operand)
		if combo < 0 {
			i = 100
		}
		if opcode == 0 {
			a = int(a / int(math.Pow(2, float64(combo))))
			i += 2
		}
		if opcode == 1 {
			b = b ^ operand
			i += 2
		}
		if opcode == 2 {
			b = combo % 8
			i += 2
		}
		if opcode == 3 {
			if a != 0 {
				i = operand
			} else {
				i += 2
			}
		}
		if opcode == 4 {
			b = b ^ c
			i += 2
		}
		if opcode == 5 {
			out = append(out, (combo % 8))
			if out[len(out)-1] != prog[len(out)-1] {
				return out
			}
			i += 2
		}
		if opcode == 6 {
			b = int(a / int(math.Pow(2, float64(combo))))
			i += 2
		}
		if opcode == 7 {
			c = int(a / int(math.Pow(2, float64(combo))))
			i += 2
		}
		//fmt.Println("Iteration:", opcode, operand, combo, a, b, c, out)
	}
	return out
}

func getOperand(operand int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return a
	} else if operand == 5 {
		return b
	} else if operand == 6 {
		return c
	} else {
		return -1
	}
}

func ChunkProgram(slice []int, chunkSize int) (chunks [][]int) {
	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}
	return append(chunks, slice)
}
