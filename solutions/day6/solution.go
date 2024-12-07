package main

import (
	"bufio"
	"fmt"
	"os"
)

type vector struct {
	x int
	y int
}

func (v vector) rotate90Right() vector {
	return vector{-v.y, v.x}
}

func (v vector) add(v2 vector) vector {
	return vector{v.x + v2.x, v.y + v2.y}
}

type guard struct {
	nCols         int
	nRows         int
	obstacles     map[vector]bool
	visited       map[vector]vector
	nVisited      int
	direction     vector
	position      vector
	initPosition  vector
	initDirection vector
}

// Returns false if went outside the lab
func (g *guard) move() (outOfBounds, loop bool) {

	//Check if outside of lab
	nextPos := g.position.add(g.direction)
	if nextPos.y < 0 || nextPos.y >= g.nRows || nextPos.x < 0 || nextPos.x >= g.nCols {
		return true, false
	}

	//Check if location was visited before in the same direction
	if direction := g.visited[nextPos]; direction == g.direction {
		return false, true
	}
	// Check if the next position collides with an obstacle
	if _, has := g.obstacles[nextPos]; has {
		g.direction = g.direction.rotate90Right()
		return false, false
	}

	// If the next move is valid
	g.position = nextPos
	if _, has := g.visited[g.position]; !has {
		g.visited[g.position] = g.direction
		g.nVisited++
	}

	//g.printLab()
	return false, false

}

func (g *guard) resetLab() {
	g.direction = g.initDirection
	g.position = g.initPosition

	g.visited = make(map[vector]vector)
	g.visited[g.position] = g.direction
	g.nVisited = 1
}

func (g *guard) printLab() {

	for row := range g.nRows {
		for col := range g.nCols {
			pos := vector{col, row}

			if g.position == pos {
				switch {
				case g.direction == vector{0, -1}:
					fmt.Print("^")
				case g.direction == vector{1, 0}:
					fmt.Print(">")
				case g.direction == vector{0, 1}:
					fmt.Print("v")
				case g.direction == vector{-1, 0}:
					fmt.Print("<")
				}
			} else if direction, has := g.visited[pos]; has {

				switch {
				case direction == vector{0, -1}:
					fmt.Print("|")
				case direction == vector{1, 0}:
					fmt.Print("-")
				case direction == vector{0, 1}:
					fmt.Print("|")
				case direction == vector{-1, 0}:
					fmt.Print("_")
				}

			} else if _, has := g.obstacles[pos]; has {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()

	}

}

func main() {
	fmt.Println("Advent of Code - Day 6") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	scanner := bufio.NewScanner(file)
	g := guard{
		nCols:     0,
		nRows:     0,
		obstacles: make(map[vector]bool),
		visited:   make(map[vector]vector),
		nVisited:  1,
		direction: vector{},
		position:  vector{},
	}
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		g.nCols = len(line)
		g.nRows++
		for col, val := range line {
			pos := vector{col, row}
			switch {
			case val == '#':
				g.obstacles[pos] = true
			case val == '^':
				g.position = vector{x: col, y: row}
				g.direction = vector{x: 0, y: -1}
				g.visited[pos] = g.direction
				g.initPosition = g.position
				g.initDirection = g.direction
			case val == '>':
				g.position = vector{x: col, y: row}
				g.direction = vector{x: 1, y: 0}
				g.visited[pos] = g.direction
				g.initPosition = g.position
				g.initDirection = g.direction
			case val == 'v':
				g.position = vector{x: col, y: row}
				g.direction = vector{x: 0, y: 1}
				g.visited[pos] = g.direction
				g.initPosition = g.position
				g.initDirection = g.direction
			case val == '<':
				g.position = vector{x: col, y: row}
				g.direction = vector{x: -1, y: 0}
				g.visited[pos] = g.direction
				g.initPosition = g.position
				g.initDirection = g.direction
			}
		}

	}
	g.printLab()
	fmt.Println()

	outOfBounds := false
	for !outOfBounds {
		outOfBounds, _ = g.move()
	}
	fmt.Printf("Part 1: %d\n", g.nVisited)
	g.resetLab()
	g.printLab()

	//--- Part 2 ---
	sum2 := 0
	total := g.nRows * g.nCols
	checked := 0
	for row := range g.nRows {
		for col := range g.nCols {
			g.resetLab()
			position := vector{col, row}
			if position == g.initPosition {
				checked++
				continue
			}

			hadObstacle := false
			if _, hadObstacle = g.obstacles[vector{col, row}]; !hadObstacle {
				g.obstacles[vector{col, row}] = true
			}

			outOfBounds, loop := false, false
			for !outOfBounds && !loop {
				outOfBounds, loop = g.move()
			}

			if loop {
				sum2++
				fmt.Println(sum2)
			}
			checked++
			fmt.Printf("(%d/%d)\n", checked, total)
			if !hadObstacle {
				delete(g.obstacles, vector{col, row})
			}
		}

	}

	fmt.Printf("Part 2: %d\n", sum2)
	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
