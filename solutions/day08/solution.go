package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day 8") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close() // Ensure the file is closed at the end

	//--- Part 1 ---
	scanner := bufio.NewScanner(file)
	towersInput := make(map[rune]map[vector]struct{})

	var rows, cols int
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for col, v := range line {
			if v != '.' {
				pos := vector{X: col, Y: rows}
				if _, has := towersInput[v]; !has {
					towersInput[v] = map[vector]struct{}{}
				}
				towersInput[v][pos] = struct{}{}
			}
		}
		rows++
	}
	fmt.Println(towersInput)
	fmt.Println(rows, cols)

	// Initialize antinodes
	antinodes := part1(towersInput, cols, rows)

	// Print antinodes
	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			pos := vector{X: col, Y: row}
			if _, has := antinodes[pos]; has {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	fmt.Printf("Part1: %d\n", len(antinodes))

	//--- Part 2 ---
	// (Part 2 logic goes here)

	antinodes2 := part2(towersInput, cols, rows)
	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			pos := vector{X: col, Y: row}
			if _, has := antinodes2[pos]; has {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	fmt.Printf("Part2: %d\n", len(antinodes2))
}

// vector represents a point in 2D space
type vector struct {
	X int
	Y int
}

func part1(allTowers map[rune]map[vector]struct{}, limX, limY int) map[vector]struct{} {

	antinodes := make(map[vector]struct{})

	for _, towers := range allTowers {
		for tower1 := range towers {
			for tower2 := range towers {
				if tower1 == tower2 {
					continue
				}

				direction := tower1.sub(tower2)
				antinodePos := tower1.add(direction)
				if antinodePos.withinBounds(limX, limY) {
					antinodes[antinodePos] = struct{}{}

				}

			}
		}

	}

	return antinodes

}

func part2(allTowers map[rune]map[vector]struct{}, limX, limY int) map[vector]struct{} {

	antinodes := make(map[vector]struct{})

	for _, towers := range allTowers {
		for tower1 := range towers {
			for tower2 := range towers {
				if tower1 == tower2 {
					continue
				}

				direction := tower1.sub(tower2)
				for i := 0; ; i++ {
					antinodePos := tower1.add(direction.scale(i))
					if !antinodePos.withinBounds(limX, limY) {
						break // Exit the loop if out of bounds
					}
					antinodes[antinodePos] = struct{}{}
				}

			}
		}

	}

	return antinodes

}

// withinBounds checks if the vector is within the grid boundaries
func (v vector) withinBounds(limX, limY int) bool {
	return v.X >= 0 && v.X < limX && v.Y >= 0 && v.Y < limY
}

// add adds two vectors
func (v1 vector) add(v2 vector) vector {
	return vector{v1.X + v2.X, v1.Y + v2.Y}
}

// Subtracts v2 from v1
func (v1 vector) sub(v2 vector) vector {
	return vector{v1.X - v2.X, v1.Y - v2.Y}
}

// Subtracts v2 from v1
func (v1 vector) scale(scalar int) vector {
	return vector{v1.X * scalar, v1.Y * scalar}
}

// check handles errors by panicking
func check(err error) {
	if err != nil {
		panic(err)
	}
}
