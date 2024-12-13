package main

import (
	"advent-of-code-2024/util"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day 12") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	scanner := bufio.NewScanner(file)
	plots := make([][]rune, 0)
	plantTypes := make(map[rune]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		plotLine := make([]rune, len(line))
		for i, c := range line {
			plotLine[i] = c
			plantTypes[c] = struct{}{}
		}
		plots = append(plots, plotLine)
		// Process each line of input
	}

	regions, cost1 := part1(plots)
	fmt.Printf("Part 1: %d\n", cost1)
	//--- Part 2 ---

	cost2 := part2(plots, regions)
	fmt.Printf("Part 2: %d\n", cost2)
	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(plots [][]rune) (regions []map[util.Vector]struct{}, cost int) {

	// Get all the regions
	totalVisited := make(map[util.Vector]struct{})
	for row := range plots {
		for col := range plots[row] {
			pos := util.Vector{X: col, Y: row}
			if _, has := totalVisited[pos]; has {
				continue
			}
			pType := plots[row][col]

			visited := make(map[util.Vector]struct{})
			regionDFS(plots, pType, pos, visited, totalVisited)
			regions = append(regions, visited)
		}

	}

	// Iterate through all positions in the region, if one is touching a different plot it has a fence
	for _, region := range regions {
		area := 0
		fence := 0
		for pos := range region {
			row := pos.Y
			col := pos.X
			pType := plots[row][col]

			// Check above
			if row+1 < len(plots) && plots[row+1][col] != pType {
				fence++
			}
			if col+1 < len(plots[0]) && plots[row][col+1] != pType {
				fence++
			}
			if row-1 >= 0 && plots[row-1][col] != pType {
				fence++
			}
			if col-1 >= 0 && plots[row][col-1] != pType {
				fence++
			}

			// Check if touching array bounds
			if row == 0 || row == len(plots)-1 {
				fence++
			}

			if col == 0 || col == len(plots[0])-1 {
				fence++
			}
			area++
		}

		cost += area * fence
	}

	return regions, cost
}

func part2(plots [][]rune, regions []map[util.Vector]struct{}) (cost int) {

	//---UNUSED IDEAS!---
	//IDEA: Get bounding box of a region, check rows, columns
	//IDEA: Edge = set of points that share the same x/y coordinate, all are next to a different plot
	//IDEA: We'll first filter the points in a region to only store those which are on the edge
	//IDEA: Count corners of a shape

	// for _, edges := range regionEdges {
	//
	// 	// Get any point from the set
	//
	// }
	for _, region := range regions {
		sideCount := sides(region)
		area := 0
		for range region {
			area++
		}
		cost += area * sideCount
	}

	return cost
}

func sides(region map[util.Vector]struct{}) (count int) {
	//---GOOD IDEAS?----
	//Calculate bounding box of region, go through all rows and columns and counts sides
	botLeft, topRight := boundingBox(region)
	//Check row by row
	for row := botLeft.Y; row <= topRight.Y; row++ {
		sidesTop, sidesBot := 0, 0

		// Count top sides
		for col := botLeft.X; col <= topRight.X; {

			length := 0
			for {
				pos := util.Vector{X: col, Y: row}
				_, has := region[pos]
				_, hasAbove := region[util.Vector{X: col, Y: row + 1}]
				if has && (!hasAbove || row == topRight.Y) {
					length++
					col++
				} else {
					col++
					break
				}

			}
			if length > 0 {
				sidesTop++
			}

		}

		// Count bot sides
		for col := botLeft.X; col <= topRight.X; {

			length := 0
			for {
				pos := util.Vector{X: col, Y: row}
				_, has := region[pos]
				_, hasBelow := region[util.Vector{X: col, Y: row - 1}]
				if has && (!hasBelow || row == botLeft.Y) {
					length++
					col++
				} else {
					col++
					break
				}

			}
			if length > 0 {
				sidesBot++
			}

		}
		count += sidesTop + sidesBot

	}

	//Check column by column
	for col := botLeft.X; col <= topRight.X; col++ {
		sidesLeft, sidesRight := 0, 0

		// Count left sides
		for row := botLeft.Y; row <= topRight.Y; {

			length := 0
			for {
				pos := util.Vector{X: col, Y: row}
				_, has := region[pos]
				_, hasLeft := region[util.Vector{X: col - 1, Y: row}]
				if has && (!hasLeft || col == botLeft.X) {
					length++
					row++
				} else {
					row++
					break
				}

			}
			if length > 0 {
				sidesLeft++
			}

		}

		// Count right sides
		for row := botLeft.Y; row <= topRight.Y; {

			length := 0
			for {
				pos := util.Vector{X: col, Y: row}
				_, has := region[pos]
				_, hasRight := region[util.Vector{X: col + 1, Y: row}]
				if has && (!hasRight || col == topRight.X) {
					length++
					row++
				} else {
					row++
					break
				}

			}
			if length > 0 {
				sidesRight++
			}

		}
		count += sidesLeft + sidesRight

	}

	return count
}

func boundingBox(region map[util.Vector]struct{}) (botLeft, topRight util.Vector) {
	for pos := range region {
		botLeft = pos
		topRight = pos
	}

	for pos := range region {
		if pos.X < botLeft.X {
			botLeft = util.Vector{X: pos.X, Y: botLeft.Y}
		}

		if pos.Y < botLeft.Y {
			botLeft = util.Vector{X: botLeft.X, Y: pos.Y}
		}

		if pos.X > topRight.X {
			topRight = util.Vector{X: pos.X, Y: topRight.Y}
		}

		if pos.Y > topRight.Y {
			topRight = util.Vector{X: topRight.X, Y: pos.Y}
		}

	}
	return botLeft, topRight
}

func printRegion(plots [][]rune, region map[util.Vector]struct{}) {
	for row := range plots {
		for col := range plots[row] {
			pos := util.Vector{X: col, Y: row}
			if _, has := region[pos]; has {
				fmt.Print(string(plots[row][col]))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkIfEdge(plots [][]rune, pos util.Vector) bool {
	row := pos.Y
	col := pos.X
	pType := plots[row][col]
	if row+1 < len(plots) && plots[row+1][col] != pType {
		return true
	}
	if col+1 < len(plots[0]) && plots[row][col+1] != pType {
		return true
	}
	if row-1 >= 0 && plots[row-1][col] != pType {
		return true
	}
	if col-1 >= 0 && plots[row][col-1] != pType {
		return true
	}

	// Check if touching array bounds
	if row == 0 || row == len(plots)-1 {
		return true
	}

	if col == 0 || col == len(plots[0])-1 {
		return true
	}

	return false

}

func regionDFS(plots [][]rune, plotType rune, pos util.Vector, visited, totalVisited map[util.Vector]struct{}) {
	if !pos.WithinBounds(len(plots[0]), len(plots)) {
		return
	}

	if _, has := visited[pos]; has {
		return
	}

	if plots[pos.Y][pos.X] != plotType {
		return
	}

	visited[pos] = struct{}{}
	totalVisited[pos] = struct{}{}

	direction := util.Vector{X: 0, Y: 1}
	above := pos.Add(direction)
	direction = direction.Rotate90Right()
	left := pos.Add(direction)
	direction = direction.Rotate90Right()
	below := pos.Add(direction)
	direction = direction.Rotate90Right()
	right := pos.Add(direction)
	direction = direction.Rotate90Right()

	regionDFS(plots, plotType, above, visited, totalVisited)
	regionDFS(plots, plotType, left, visited, totalVisited)
	regionDFS(plots, plotType, below, visited, totalVisited)
	regionDFS(plots, plotType, right, visited, totalVisited)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
