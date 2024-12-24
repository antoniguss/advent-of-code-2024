package main

import (
	"bufio"
	"fmt"
	"github.com/Antyhot/advent-of-code-24/util"
	"github.com/Antyhot/advent-of-code-24/util/queue"
	"os"
	"time"
)

const (
	inputFile = "input.txt"
)

func main() {
	fmt.Println("Advent of Code - Day 20") // Placeholder for day number

	start, end, maze := getInput()

	//--- Part 1 ---
	startTime1 := time.Now() // Start time for Part 1
	var sol1 int = part1(start, end, maze)
	elapsedTime1 := time.Since(startTime1) // Calculate elapsed time for Part 1
	fmt.Printf("Part1: %d (Time: %s)\n", sol1, elapsedTime1)

	startTime2 := time.Now() // Start time for Part 2
	var sol2 int = part2(start, end, maze)
	elapsedTime2 := time.Since(startTime2) // Calculate elapsed time for Part 2
	fmt.Printf("Part2: %d (Time: %s)\n", sol2, elapsedTime2)

}

func getInput() (util.Vector, util.Vector, [][]bool) {
	file, err := os.Open(inputFile)
	check(err)
	scanner := bufio.NewScanner(file)

	var start util.Vector
	var end util.Vector

	maze := [][]bool{}

	rows := 0
	for scanner.Scan() {
		line := scanner.Text()
		mazeLine := make([]bool, len(line))
		for i, c := range line {
			switch c {
			case '#':
				mazeLine[i] = true
			case 'S':
				{
					start = util.Vector{X: i, Y: rows}
				}
			case 'E':
				{
					end = util.Vector{X: i, Y: rows}
				}
			}
		}
		maze = append(maze, mazeLine)
		rows++
	}
	return start, end, maze
}

type maze [][]bool

func (m maze) get(pos util.Vector) bool {
	if pos.WithinBounds(len(m[0]), len(m)) {
		return m[pos.Y][pos.X]
	}
	return true
}

func calcDistances(m maze, end, start util.Vector) (distance map[util.Vector]int) {
	distance = make(map[util.Vector]int)
	cur := start
	distance[cur] = 0
	var prev util.Vector

	dist := 0
	for cur != end {
		dist++
		options := []util.Vector{cur.Above(1), cur.Below(1), cur.Left(1), cur.Right(1)}
		for _, opt := range options {
			if !m.get(opt) && opt != prev {
				prev = cur
				cur = opt
				distance[opt] = dist
				break
			}
		}
	}
	return distance
}

type pair struct {
	util.Vector
	distance int
}

// Solution using BFS, worked in ~2s
// Could be better if we only examined points on the path (since the start and end of a cheat has to be on the path)
func part2(start, end util.Vector, m maze) (cheatCount int) {
	//fmt.Println(start, end)

	distance := calcDistances(m, end, start)

	cur := start
	prev := util.Vector{}
	for cur != end {
		curDist, _ := distance[cur]

		maxSkipLen := 20
		cheatThresh := 100
		//Start a BFS search to check all positions within 20 picoseconds
		checked := map[util.Vector]struct{}{}

		queue := queue.NewQueue()
		queue.Enqueue(pair{
			cur,
			0,
		})
		for !queue.IsEmpty() {
			cheat := queue.Dequeue().(pair)

			if _, has := checked[cheat.Vector]; has {
				continue
			}

			if cheatDist, has := distance[cheat.Vector]; has && cheatDist > curDist {
				skip := cheatDist - curDist - cheat.distance
				if skip >= cheatThresh {
					cheatCount++
				}

			}
			checked[cheat.Vector] = struct{}{}

			//Add other possible options
			options := []pair{
				{cheat.Vector.Above(1), cheat.distance + 1},
				{cheat.Vector.Right(1), cheat.distance + 1},
				{cheat.Vector.Below(1), cheat.distance + 1},
				{cheat.Vector.Left(1), cheat.distance + 1},
			}

			for _, opt := range options {
				if opt.distance > maxSkipLen {
					continue
				}

				if _, has := checked[opt.Vector]; has {
					continue
				}
				queue.Enqueue(opt)
			}
		}

		//Get the next step
		options := []util.Vector{cur.Above(1), cur.Below(1), cur.Left(1), cur.Right(1)}
		for _, opt := range options {
			if !m.get(opt) && opt != prev {
				prev = cur
				cur = opt
				break
			}
		}
	}

	return cheatCount
}

func part1(start, end util.Vector, m maze) (cheatCount int) {
	//fmt.Println(start, end)

	distance := calcDistances(m, end, start)

	cur := start
	var prev util.Vector

	for cur != end {
		curDist, _ := distance[cur]
		//Check above
		if m.get(cur.Above(1)) && !m.get(cur.Above(2)) {
			cheatDist, has := distance[cur.Above(2)]
			if has && cheatDist > curDist {
				skip := cheatDist - curDist - 2
				if skip >= 100 {
					cheatCount++
				}

			}

		}

		//Check right
		if m.get(cur.Right(1)) && !m.get(cur.Right(2)) {
			cheatDist, has := distance[cur.Right(2)]

			if has && cheatDist > curDist {
				skip := cheatDist - curDist - 2
				if skip >= 100 {
					cheatCount++
				}
			}
		}

		//Check below
		if m.get(cur.Below(1)) && !m.get(cur.Below(2)) {
			cheatDist, has := distance[cur.Below(2)]

			if has && cheatDist > curDist {
				skip := cheatDist - curDist - 2
				if skip >= 100 {
					cheatCount++
				}
			}
		}

		//Check left
		if m.get(cur.Left(1)) && !m.get(cur.Left(2)) {
			cheatDist, has := distance[cur.Left(2)]
			if has && cheatDist > curDist {
				skip := cheatDist - curDist - 2
				if skip >= 100 {
					cheatCount++
				}
			}
		}

		options := []util.Vector{cur.Above(1), cur.Below(1), cur.Left(1), cur.Right(1)}
		for _, opt := range options {
			if !m.get(opt) && opt != prev {
				prev = cur
				cur = opt
				break
			}
		}
	}

	return cheatCount
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
