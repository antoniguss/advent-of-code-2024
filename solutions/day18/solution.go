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
	grid := make([][]bool, rows)
	for i := range cols {
		grid[i] = make([]bool, cols)
	}

	scanner := bufio.NewScanner(file)
	for i := 0; i < first && scanner.Scan(); i++ {
		line := strings.Split(scanner.Text(), ",")
		col, _ := strconv.Atoi(line[0])
		row, _ := strconv.Atoi(line[1])
		grid[row][col] = true
	}

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
	sol1 := part1(rows, cols, grid)
	fmt.Printf("Part1: %d\n", sol1)
	//--- Part 2 ---

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(rows, cols int, grid [][]bool) (cost int) {
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

func part2(rows, cols int, grid [][]bool) (X, Y int) {

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
