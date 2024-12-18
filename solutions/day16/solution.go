package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Antyhot/advent-of-code-24/util"
	"github.com/Antyhot/advent-of-code-24/util/heap"
	"github.com/Antyhot/advent-of-code-24/util/queue"
)

type move struct {
	pos    util.Vector
	facing util.Vector
}

type costMove struct {
	move
	cost int
}

func moveComparator(a *costMove, b *costMove) bool {
	return a.cost < b.cost
}

func main() {
	fmt.Println("Advent of Code - Day 16") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---

	scanner := bufio.NewScanner(file)
	maze := make([][]rune, 0)
	rows := 0

	var sPos util.Vector
	var ePos util.Vector

	for scanner.Scan() {
		line := scanner.Text()
		mazeLine := make([]rune, len(line))
		for i, c := range line {
			if c == 'S' {
				sPos = util.Vector{X: i, Y: rows}
			} else if c == 'E' {
				ePos = util.Vector{X: i, Y: rows}
			} else {
				mazeLine[i] = c
			}

		}

		maze = append(maze, mazeLine)
		// Process each line of input
		fmt.Println(line)
		rows++
	}
	fmt.Println(sPos)
	fmt.Println(ePos)

	cost1 := part1(maze, sPos, ePos)
	fmt.Printf("Part1: %d\n", cost1)

	//--- Part 2 ---
	cost2 := part2(maze, sPos, ePos)
	fmt.Printf("Part2: %d\n", cost2)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(maze [][]rune, sPos, ePos util.Vector) (cost int) {

	moveQueue := heap.NewHeap(moveComparator)

	startMove := move{
		pos:    sPos,
		facing: util.Vector{X: 1, Y: 0},
	}

	start := costMove{
		cost: 0,
		move: startMove,
	}
	moveQueue.Push(start)
	seen := map[move]struct{}{}
	seen[startMove] = struct{}{}

	for !moveQueue.IsEmpty() {
		cur, _ := moveQueue.Pop()
		// fmt.Println(moveQueue.GetSize())
		// fmt.Printf("cur: %v\n", cur)
		// fmt.Printf("moveQueue.GetSize(): %v\n", moveQueue.GetSize())
		// fmt.Printf("moveQueue: %v\n", moveQueue)

		if cur.pos == ePos {
			return cur.cost
		}
		seen[cur.move] = struct{}{}

		option1 := costMove{
			move: move{
				pos:    cur.pos.Add(cur.facing),
				facing: cur.facing,
			},
			cost: cur.cost + 1,
		}

		option2 := costMove{
			move: move{
				pos:    cur.pos,
				facing: cur.facing.Rotate90Right(),
			},
			cost: cur.cost + 1000,
		}

		option3 := costMove{
			move: move{
				pos:    cur.pos,
				facing: cur.facing.Rotate90Left(),
			},
			cost: cur.cost + 1000,
		}
		options := []costMove{option1, option2, option3}

		for _, option := range options {
			if maze[option.pos.Y][option.pos.X] == '#' {
				// fmt.Printf("option rejected:running into wall: %v\n", option)
				continue
			}

			if _, has := seen[option.move]; has {
				// fmt.Printf("option rejected : already visited: %v\n", option)
				continue
			}

			// fmt.Printf("adding option: %v\n", option)
			moveQueue.Push(option)

		}

	}
	return cost
}

func part2(maze [][]rune, sPos, ePos util.Vector) (cost int) {

	moveQueue := heap.NewHeap(moveComparator)

	startMove := move{
		pos:    sPos,
		facing: util.Vector{X: 1, Y: 0},
	}

	start := costMove{
		cost: 0,
		move: startMove,
	}
	moveQueue.Push(start)
	//Instead of seen we track the position and it's lowest cost
	lowest := map[move]int{}
	lowest[startMove] = 0

	//Tracks a set of previous moves for the current set
	backtrack := map[move](map[move]struct{}){}

	// Max int value
	bestCost := int(^uint(0) >> 1)
	finalStates := map[move]struct{}{}
	for !moveQueue.IsEmpty() {
		cur, _ := moveQueue.Pop()
		if lowestCost, visited := lowest[cur.move]; visited {
			if cur.cost > lowestCost {
				continue
			}

		}

		lowest[cur.move] = cur.cost

		//We got to the final position in the lowest cost
		if cur.pos == ePos {
			if cur.cost > bestCost {
				finalStates[cur.move] = struct{}{}
				break

			}
			bestCost = cur.cost

		}

		option1 := costMove{
			move: move{
				pos:    cur.pos.Add(cur.facing),
				facing: cur.facing,
			},
			cost: cur.cost + 1,
		}

		option2 := costMove{
			move: move{
				pos:    cur.pos,
				facing: cur.facing.Rotate90Right(),
			},
			cost: cur.cost + 1000,
		}

		option3 := costMove{
			move: move{
				pos:    cur.pos,
				facing: cur.facing.Rotate90Left(),
			},
			cost: cur.cost + 1000,
		}
		options := []costMove{option1, option2, option3}

		for _, option := range options {
			if maze[option.pos.Y][option.pos.X] == '#' {
				// fmt.Printf("option rejected:running into wall: %v\n", option)
				continue
			}

			if lowestCost, visited := lowest[option.move]; visited {
				if option.cost > lowestCost {
					continue
				} else if option.cost < lowestCost {
					backtrack[option.move] = make(map[move]struct{})
					lowest[option.move] = option.cost
				}
			} else {
				lowest[option.move] = option.cost
				backtrack[option.move] = make(map[move]struct{})
			}
			// fmt.Printf("%v -> %v\n", cur.move, option.move)
			backtrack[option.move][cur.move] = struct{}{}

			// fmt.Printf("adding option: %v\n", option)
			moveQueue.Push(option)
		}
	}

	seen := map[move]struct{}{}
	queue := queue.NewQueue()

	for state := range finalStates {
		seen[state] = struct{}{}
		queue.Enqueue(state)
	}

	for !queue.IsEmpty() {
		state := queue.Dequeue().(move)
		for prevState := range backtrack[state] {
			if _, has := seen[prevState]; has {
				continue
			}

			seen[prevState] = struct{}{}
			queue.Enqueue(prevState)
		}

	}

	seenPos := map[util.Vector]struct{}{}
	for state := range seen {
		seenPos[state.pos] = struct{}{}
	}

	// for row := range len(maze) {
	// 	for col := range len(maze[row]) {
	// 		pos := util.Vector{X: col, Y: row}
	//
	// 		if _, has := seenPos[pos]; has {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Print(string(maze[row][col]))
	// 		}
	//
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Printf("seenPos: %v\n", seenPos)

	return len(seenPos)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
