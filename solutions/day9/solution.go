package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day 9") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)

	disc := make([]int, 0)
	for i, c := range line {
		n := int(c - '0')
		block := make([]int, n)
		if i%2 == 0 {
			//reading file block
			for v := 0; v < n; v++ {
				block[v] = i / 2
			}
		} else {
			for v := 0; v < n; v++ {
				block[v] = -1
			}
			//reading empty space
		}
		disc = append(disc, block...)

	}
	// printDisc(disc)
	disc1 := part1(disc)
	// printDisc(disc1)
	fmt.Printf("Part1: %v\n", calcChecksum(disc1))

	//--- Part 2 ---
	printDisc(disc)
	disc2 := part2(disc)
	printDisc(disc2)

	fmt.Printf("Part2: %v\n", calcChecksum(disc2))
	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func printDisc(disc []int) {
	for _, val := range disc {
		if val == -1 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", val)
		}
	}

	fmt.Println()
}

func part1(input []int) []int {
	disc := make([]int, len(input))
	copy(disc, input)
	l := 0
	r := len(disc) - 1

	// While l < r
	for {
		//Move l to next free block
		for disc[l] > -1 {
			l++
		}

		for disc[r] == -1 {
			r--
		}

		if l >= r {
			break
		}

		disc[l], disc[r] = disc[r], disc[l]
	}

	return disc

}

func part2(input []int) []int {

	disc := make([]int, len(input))
	copy(disc, input)
	r := len(disc) - 1

	// While l < r
	for r > 0 {
		//Move right pointer to end of rightmost block
		for disc[r] == -1 {
			r--
		}
		id := disc[r]

		//Calculate size of block
		sizeBlock := 0
		for r-sizeBlock >= 0 && disc[r-sizeBlock] == id {
			sizeBlock++
		}

		if r == 0 {
			break
		}

		sizeEmpty := 0
		l := 0
		fmt.Printf("r: %v\n", r)
		for {

			//Find first empty block
			for disc[l] > -1 {
				l++
			}
			if l >= r {
				break
			}
			//Calculate size of empty block
			sizeEmpty = 0
			for disc[l+sizeEmpty] == -1 {
				sizeEmpty++
			}

			if sizeEmpty >= sizeBlock {
				break
			}

			l += sizeEmpty

		}

		if sizeEmpty < sizeBlock {
			r -= sizeBlock
			continue
		}

		for i := range sizeBlock {
			disc[l+i] = id
			disc[r-i] = -1

		}

	}

	return disc

}
func calcChecksum(disc []int) int {
	checksum := 0
	for i, val := range disc {
		if val == -1 {
			continue
		}

		checksum += i * val
	}

	return checksum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
