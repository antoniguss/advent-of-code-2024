package main

import (
	"advent-of-code-2024/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day 10") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	//IDEA: recursively search the 4 (or less) possible paths to return a list of reachable points

	topoMap := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mapRow := make([]int, len(line))
		for i, c := range line {
			mapRow[i] = int(c - '0')

		}
		topoMap = append(topoMap, mapRow)

		// Process each line of input
	}
	sum1 := part1(topoMap)
	fmt.Printf("Part 1: %d\n", sum1)

	//--- Part 2 ---
	sum2 := part2(topoMap)
	fmt.Printf("Part 2: %d\n", sum2)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(topoMap [][]int) int {

	var sum int = 0
	for row := range topoMap {
		for col := range topoMap[row] {
			if topoMap[row][col] == 0 {
				reachable := make(map[util.Vector]struct{}, 0)
				findPeaks(topoMap, util.Vector{X: col, Y: row}, -1, reachable)
				sum += len(reachable)

			}

		}

	}

	return sum
}

func part2(topoMap [][]int) int {

	var sum int = 0
	for row := range topoMap {
		for col := range topoMap[row] {
			if topoMap[row][col] == 0 {
				grade := findTrails(topoMap, util.Vector{X: col, Y: row}, -1)
				sum += grade

			}

		}

	}

	return sum
}

func findPeaks(topoMap [][]int, pos util.Vector, prev int, peakSet map[util.Vector]struct{}) {
	rows := len(topoMap)
	cols := len(topoMap[0])

	if !pos.WithinBounds(cols, rows) {
		return
	}

	val := topoMap[pos.Y][pos.X]

	if val-prev != 1 {
		return
	}

	if val == 9 {
		peakSet[pos] = struct{}{}
		return
	}

	direction := util.Vector{X: 0, Y: 1}
	above := pos.Add(direction)
	direction = direction.Rotate90Right()
	left := pos.Add(direction)
	direction = direction.Rotate90Right()
	below := pos.Add(direction)
	direction = direction.Rotate90Right()
	right := pos.Add(direction)
	direction = direction.Rotate90Right()

	findPeaks(topoMap, above, val, peakSet)
	findPeaks(topoMap, left, val, peakSet)
	findPeaks(topoMap, below, val, peakSet)
	findPeaks(topoMap, right, val, peakSet)
}

func findTrails(topoMap [][]int, pos util.Vector, prev int) int {
	rows := len(topoMap)
	cols := len(topoMap[0])

	if !pos.WithinBounds(cols, rows) {
		return 0
	}

	val := topoMap[pos.Y][pos.X]

	if val-prev != 1 {
		return 0
	}

	if val == 9 {
		return 1
	}

	direction := util.Vector{X: 0, Y: 1}
	above := pos.Add(direction)
	direction = direction.Rotate90Right()
	left := pos.Add(direction)
	direction = direction.Rotate90Right()
	below := pos.Add(direction)
	direction = direction.Rotate90Right()
	right := pos.Add(direction)
	direction = direction.Rotate90Right()

	return findTrails(topoMap, above, val) + findTrails(topoMap, left, val) + findTrails(topoMap, below, val) + findTrails(topoMap, right, val)
}

// func part1(topoMap[][])

func check(e error) {
	if e != nil {
		panic(e)
	}
}
