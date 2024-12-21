package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Antyhot/advent-of-code-24/util"
	"github.com/Antyhot/advent-of-code-24/util/queue"
)

func main() {
	fmt.Println("Advent of Code - Day 18") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	rows := 71
	cols := 71
	first := 1024
	coords := make([]util.Vector, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		col, _ := strconv.Atoi(line[0])
		row, _ := strconv.Atoi(line[1])
		coords = append(coords, util.Vector{X: col, Y: row})
	}

	sol1 := part1(rows, cols, first, coords)
	fmt.Printf("Part1: %d\n", sol1)
	//--- Part 2 ---
	X, Y := part2(rows, cols, coords)
	fmt.Printf("Part2: %d,%d\n", X, Y)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func displayGrid(rows, cols, int, grid [][]bool) {

	for row := range rows {
		for col := range cols {
			if grid[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(rows, cols, first int, coords []util.Vector) (cost int) {

	grid := make([][]bool, rows)
	for i := range cols {
		grid[i] = make([]bool, cols)
	}

	for i := 0; i < first && i < len(coords); i++ {
		col, row := coords[i].Get()
		grid[row][col] = true
	}

	pq := queue.NewQueue()

	start := move{
		util.Vector{
			X: 0,
			Y: 0,
		},
		0,
	}

	endPos := util.Vector{X: rows - 1, Y: cols - 1}
	seen := make(map[util.Vector]struct{})
	pq.Enqueue(start)

	for !pq.IsEmpty() {
		cur := pq.Dequeue().(move)
		seen[cur.Vector] = struct{}{}
		// fmt.Printf("cur: %v\n", cur)

		if cur.Vector == endPos {
			return cur.cost
		}

		moveUp := move{cur.Above(1), cur.cost + 1}
		moveDown := move{cur.Below(1), cur.cost + 1}
		moveLeft := move{cur.Left(1), cur.cost + 1}
		moveRight := move{cur.Right(1), cur.cost + 1}

		moveOptions := []move{moveUp, moveDown, moveLeft, moveRight}

		for _, move := range moveOptions {
			if !move.Vector.WithinBounds(rows, cols) {
				continue
			}
			if grid[move.Y][move.X] {
				continue
			}

			if _, has := seen[move.Vector]; has {
				continue
			}

			seen[move.Vector] = struct{}{}
			pq.Enqueue(move)
		}

	}
	return -1
}

func bfs(rows, cols int, grid [][]bool) (cost int) {
	// pq := heap.NewHeap(moveComparator)
	pq := queue.NewQueue()

	start := move{
		util.Vector{
			X: 0,
			Y: 0,
		},
		0,
	}

	endPos := util.Vector{X: rows - 1, Y: cols - 1}
	seen := make(map[util.Vector]struct{})
	pq.Enqueue(start)

	for !pq.IsEmpty() {
		cur := pq.Dequeue().(move)
		seen[cur.Vector] = struct{}{}
		// fmt.Printf("cur: %v\n", cur)

		if cur.Vector == endPos {
			return cur.cost
		}

		moveUp := move{cur.Above(1), cur.cost + 1}
		moveDown := move{cur.Below(1), cur.cost + 1}
		moveLeft := move{cur.Left(1), cur.cost + 1}
		moveRight := move{cur.Right(1), cur.cost + 1}

		moveOptions := []move{moveUp, moveDown, moveLeft, moveRight}

		for _, move := range moveOptions {
			if !move.Vector.WithinBounds(rows, cols) {
				continue
			}
			if grid[move.Y][move.X] {
				continue
			}

			if _, has := seen[move.Vector]; has {
				continue
			}

			seen[move.Vector] = struct{}{}
			pq.Enqueue(move)
		}

	}
	return -1

}

func part2(rows, cols int, coords []util.Vector) (X, Y int) {

	grid := make([][]bool, rows)
	for i := range cols {
		grid[i] = make([]bool, cols)
	}

	//Last possible path after 1024 bytes fall
	for i := 0; i < len(coords); i++ {
		col, row := coords[i].Get()
		grid[row][col] = true

		if bfs(rows, cols, grid) == -1 {
			return col, row
		}

	}

	return X, Y
}

type move struct {
	util.Vector
	cost int
}

func moveComparator(a *move, b *move) bool {
	return a.cost < b.cost
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
