package main

import (
	"advent-of-code-2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	id int
	p  util.Vector
	v  util.Vector
}

func main() {
	fmt.Println("Advent of Code - Day 14") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	cols := 101
	rows := 103

	scanner := bufio.NewScanner(file)
	robots := make([]robot, 0)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		cords := strings.Split(parts[0][2:], ",")
		pos := util.Vector{}
		pos.X, _ = strconv.Atoi(cords[0])
		pos.Y, _ = strconv.Atoi(cords[1])

		velocity := strings.Split(parts[1][2:], ",")
		vel := util.Vector{}
		vel.X, _ = strconv.Atoi(velocity[0])
		vel.Y, _ = strconv.Atoi(velocity[1])

		newRobot := robot{id: i, p: pos, v: vel}
		robots = append(robots, newRobot)
		// Process each line of input
	}

	score1 := part1(robots, cols, rows, 100)
	fmt.Printf("Part1: %d\n", score1)

	//--- Part 2 ---
	score2 := part2(robots, cols, rows, 99999)
	fmt.Printf("Part2: %d\n", score2)
	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func (r *robot) move(cols, rows, moves int) {
	newPos := r.p.Add(r.v.Scale(moves))
	newPos.X = mod(newPos.X, cols)
	newPos.Y = mod(newPos.Y, rows)

	r.p = newPos

}

func mod(a, b int) int {
	return (a%b + b) % b

}

func part1(robots []robot, cols, rows, moves int) (score int) {

	leftTopCount := 0
	leftBotCount := 0
	rightTopCount := 0
	rightBotCount := 0

	// leftStart := 0
	leftEnd := (cols - 2) / 2

	rightStart := (cols + 1) / 2
	// rightStop := cols

	// topStart := 0
	topEnd := (rows - 2) / 2

	botStart := (rows + 1) / 2
	// botEnd := rows - 1

	for _, robot := range robots {
		robot.move(cols, rows, moves)

		if robot.p.X <= leftEnd {
			if robot.p.Y <= topEnd {
				leftTopCount++
				continue
			}

			if robot.p.Y >= botStart {
				leftBotCount++
				continue
			}

		}

		if robot.p.X >= rightStart {
			if robot.p.Y <= topEnd {
				rightTopCount++
				continue
			}

			if robot.p.Y >= botStart {
				rightBotCount++
				continue
			}

		}

	}

	return leftTopCount * leftBotCount * rightTopCount * rightBotCount
}
func part2(robots []robot, cols, rows, moves int) (score int) {

	for i := range moves {
		for j := range robots {
			robots[j].move(cols, rows, 1)
		}
		fmt.Println(i)
		if i >= 7500 {
			displayImage(robots, cols, rows)
			fmt.Printf("----%d----\n", i+1)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
		}
		fmt.Print("\033[H\033[2J")
	}

	return score
}

func displayImage(robots []robot, cols, rows int) {
	pixels := make([][]bool, rows)
	for i := range rows {
		pixels[i] = make([]bool, cols)
	}

	for _, robot := range robots {
		pixels[robot.p.Y][robot.p.X] = true

	}

	for row := range rows {
		for col := range cols {
			if pixels[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
