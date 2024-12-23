package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Antyhot/advent-of-code-24/util"
)

func main() {
	fmt.Println("Advent of Code - Day 20") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	var start util.Vector
	var end util.Vector

	maze := [][]bool{}

	rows := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
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

	//--- Part 1 ---
	var sol1 int = part1(start, end, maze)
	fmt.Printf("Part1: %d\n", sol1)
	//--- Part 2 ---
	var sol2 int = part2()
	fmt.Printf("Part2: %d\n", sol2)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

type maze [][]bool

func (m maze) get(pos util.Vector) bool {
	if pos.WithinBounds(len(m[0]), len(m)) {
		return m[pos.Y][pos.X]
	}
	return true
}

func part1(start, end util.Vector, m maze) (cheatCount int) {
	//fmt.Println(start, end)
	distance := make(map[util.Vector]int)

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

	cur = start

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

func part2() (cheats int) {
	return cheats
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
