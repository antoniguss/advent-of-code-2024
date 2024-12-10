package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Advent of Code - Day 4") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	mapping := map[rune]int{
		'X': 1,
		'M': 2,
		'A': 4,
		'S': 8,
	}
	revMapping := map[int]rune{
		1: 'X',
		2: 'M',
		4: 'A',
		8: 'S',
	}

	scanner := bufio.NewScanner(file)

	var chars [][]int
	for scanner.Scan() {

		line := scanner.Text()
		row := make([]int, len(line))

		for i, char := range line {
			row[i] = mapping[char]
		}
		chars = append(chars, row)

	}
	used := make([][]bool, len(chars))
	for i := range len(chars) {
		used[i] = make([]bool, len(chars[0]))
	}
	//char[i][j] means i-th row, j-th column

	//--- Part 1 ---
	//Find string "XMAS" horizontal, vertical, diagonal, backwards,
	nCols := len(chars[0])
	nRows := len(chars)
	allowed := 15 //4 characters will only sum up to 15 if they consist of 'X','M','A','S' (in any order)
	found := 0
	// 1) Horizontal
	for row := range nRows {
		for col := 0; col <= nCols-4; col++ {
			//Check if the characters sum up to 'allowed'
			if chars[row][col]+chars[row][col+1]+chars[row][col+2]+chars[row][col+3] != allowed {
				continue
			}

			//Check if we have 'XMAS'
			if chars[row][col] < chars[row][col+1] &&
				chars[row][col+1] < chars[row][col+2] &&
				chars[row][col+2] < chars[row][col+3] {
				found++
				used[row][col] = true
				used[row][col+1] = true
				used[row][col+2] = true
				used[row][col+3] = true
			}
			//Also check backwards
			if chars[row][col] > chars[row][col+1] &&
				chars[row][col+1] > chars[row][col+2] &&
				chars[row][col+2] > chars[row][col+3] {
				found++
				used[row][col] = true
				used[row][col+1] = true
				used[row][col+2] = true
				used[row][col+3] = true
			}
		}
	}

	// 2) Vertical
	for col := range nCols {
		for row := 0; row <= nRows-4; row++ {
			//Check if the characters sum up to 'allowed'
			if chars[row][col]+chars[row+1][col]+chars[row+2][col]+chars[row+3][col] != allowed {
				continue
			}

			//Check if we have 'XMAS'
			if chars[row][col] < chars[row+1][col] &&
				chars[row+1][col] < chars[row+2][col] &&
				chars[row+2][col] < chars[row+3][col] {
				found++
				used[row][col] = true
				used[row+1][col] = true
				used[row+2][col] = true
				used[row+3][col] = true
				continue
			}
			//If not, check backwards
			if chars[row][col] > chars[row+1][col] &&
				chars[row+1][col] > chars[row+2][col] &&
				chars[row+2][col] > chars[row+3][col] {
				found++
				used[row][col] = true
				used[row+1][col] = true
				used[row+2][col] = true
				used[row+3][col] = true
			}
		}
	}

	// 3) Diagonal \
	for col := 0; col <= nCols-4; col++ {
		for row := 0; row <= nRows-4; row++ {
			//Check if the characters sum up to 'allowed'
			if chars[row][col]+chars[row+1][col+1]+chars[row+2][col+2]+chars[row+3][col+3] != allowed {
				continue
			}

			//Check if we have 'XMAS'
			if chars[row][col] < chars[row+1][col+1] &&
				chars[row+1][col+1] < chars[row+2][col+2] &&
				chars[row+2][col+2] < chars[row+3][col+3] {
				found++
				used[row][col] = true
				used[row+1][col+1] = true
				used[row+2][col+2] = true
				used[row+3][col+3] = true
				continue
			}
			//If not, check backwards
			if chars[row][col] > chars[row+1][col+1] &&
				chars[row+1][col+1] > chars[row+2][col+2] &&
				chars[row+2][col+2] > chars[row+3][col+3] {
				found++
				used[row][col] = true
				used[row+1][col+1] = true
				used[row+2][col+2] = true
				used[row+3][col+3] = true
			}
		}
	}
	// 4) Diagonal /
	for col := 0; col <= nCols-4; col++ {
		for row := 3; row < nRows; row++ {
			//Check if the characters sum up to 'allowed'
			if chars[row][col]+chars[row-1][col+1]+chars[row-2][col+2]+chars[row-3][col+3] != allowed {
				continue
			}

			//Check if we have 'XMAS'
			if chars[row][col] < chars[row-1][col+1] &&
				chars[row-1][col+1] < chars[row-2][col+2] &&
				chars[row-2][col+2] < chars[row-3][col+3] {
				found++
				used[row][col] = true
				used[row-1][col+1] = true
				used[row-2][col+2] = true
				used[row-3][col+3] = true
				continue
			}
			//If not, check backwards
			if chars[row][col] > chars[row-1][col+1] &&
				chars[row-1][col+1] > chars[row-2][col+2] &&
				chars[row-2][col+2] > chars[row-3][col+3] {
				found++
				used[row][col] = true
				used[row-1][col+1] = true
				used[row-2][col+2] = true
				used[row-3][col+3] = true
			}
		}
	}

	//Print characters that were used
	for row := range used {
		for col := range used[row] {
			fmt.Print(string(revMapping[chars[row][col]]))
		}
		fmt.Println()
	}
	fmt.Println("-----------------------------")

	for row := range used {
		for col := range used[row] {
			if used[row][col] {
				fmt.Print(string(revMapping[chars[row][col]]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	end := time.Now()
	fmt.Printf("Found the word 'XMAS': %d, took %v\n", found, end.Sub(start))

	//--- Part 2 ---
	//the mapping approach turned out to be a bit unnecessary for this part
	start2 := time.Now()
	M := mapping['M']
	A := mapping['A']
	S := mapping['S']

	used2 := make([][]bool, len(chars))
	for i := range len(chars) {
		used2[i] = make([]bool, len(chars[0]))
	}
	found2 := 0
	for row := 0; row <= nRows-3; row++ {
		for col := 0; col <= nCols-3; col++ {

			if chars[row+1][col+1] != A {
				continue
			}
			used2[row+1][col+1] = true

			//Check diagonal \
			checked1 := false
			if chars[row][col] == M &&
				chars[row+2][col+2] == S {
				checked1 = true
				used2[row][col] = true
				used2[row+2][col+2] = true
			} else if chars[row][col] == S &&
				chars[row+2][col+2] == M {
				checked1 = true
				used2[row][col] = true
				used2[row+2][col+2] = true
			}

			//Check diagonal /
			checked2 := false
			if chars[row][col+2] == M &&
				chars[row+2][col] == S {
				checked2 = true
				used2[row][col+2] = true
				used2[row+2][col] = true
			} else if chars[row][col+2] == S &&
				chars[row+2][col] == M {
				checked2 = true
				used2[row][col+2] = true
				used2[row+2][col] = true
			}

			if checked1 && checked2 {
				found2++
			}
		}
	}
	fmt.Println("-----------------------------")

	for row := range used2 {
		for col := range used2[row] {
			if used2[row][col] {
				fmt.Print(string(revMapping[chars[row][col]]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	end2 := time.Now()
	fmt.Printf("Found X-MASes %d, took %v\n", found2, end2.Sub(start2))

	//--- Cleanup
	err = file.Close()
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
