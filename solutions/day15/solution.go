package main

import (
	"advent-of-code-2024/util"
	"bufio"
	"fmt"
	"os"
)

type object struct {
	pos       util.Vector
	isMovable bool
	isRobot   bool
}

func main() {
	fmt.Println("Advent of Code - Day 15") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	rows := 0
	cols := len(line)
	objects := make(map[util.Vector]object, 0)
	var robot object
	for len(line) != 0 && line[0] == '#' {
		// Process each line of input
		for col, c := range line {
			pos := util.Vector{X: col, Y: rows}

			switch c {
			case '#':
				objects[pos] = object{pos: pos, isMovable: false, isRobot: false}
			case 'O':
				objects[pos] = object{pos: pos, isMovable: true, isRobot: false}
			case '@':
				robot = object{pos: pos, isMovable: true, isRobot: true}
				objects[pos] = robot
			}

		}

		scanner.Scan()
		line = scanner.Text()
		rows++
	}
	//Load instructions
	instructions := make([]util.Vector, 0)
	for scanner.Scan() {
		instructionLine := scanner.Text()
		for _, c := range instructionLine {
			switch {
			case c == '^':
				instructions = append(instructions, util.Vector{X: 0, Y: -1})
			case c == '>':
				instructions = append(instructions, util.Vector{X: 1, Y: 0})
			case c == 'v':
				instructions = append(instructions, util.Vector{X: 0, Y: 1})
			case c == '<':
				instructions = append(instructions, util.Vector{X: -1, Y: 0})
			}
		}
	}
	sum1 := part1(objects, instructions, robot, rows, cols)
	fmt.Printf("Part1: %d\n", sum1)

	//--- Part 2 ---

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(objects map[util.Vector]object, instruction []util.Vector, robot object, rows, cols int) (sum int) {
	// fmt.Println("Initial state:")
	// printObjects(objects, rows, cols)
	for _, instruction := range instruction {
		// switch {
		// case instruction.X == 0 && instruction.Y == -1:
		// 	fmt.Println("Move ^")
		// case instruction.X == 1 && instruction.Y == 0:
		// 	fmt.Println("Move >")
		// case instruction.X == 0 && instruction.Y == 1:
		// 	fmt.Println("Move v")
		// case instruction.X == -1 && instruction.Y == 0:
		// 	fmt.Println("Move <")
		// }
		robot.move(objects, instruction)

		// printObjects(objects, rows, cols)

	}

	for _, object := range objects {

		if object.isMovable && !object.isRobot {
			sum += 100*object.pos.Y + object.pos.X
		}
	}
	return sum
}

func printObjects(objects map[util.Vector]object, rows, cols int) {
	for row := range rows {
		for col := range cols {
			pos := util.Vector{col, row}
			object, has := objects[pos]
			if has {
				if object.isRobot {
					fmt.Print("@")
					continue
				}

				if object.isMovable {
					fmt.Print("O")
					continue
				}

				fmt.Print("#")
				continue
			}

			fmt.Print(".")

		}

		fmt.Println()
	}

	fmt.Println()
}

func (o *object) move(objects map[util.Vector]object, instruction util.Vector) bool {

	if !o.isMovable {
		return false
	}
	//Before attempting to move check if there's an object in front. If yes, ask them to move
	nextPos := o.pos.Add(instruction)
	if inFront, has := objects[nextPos]; has {
		if !inFront.move(objects, instruction) {
			return false
		}

	}

	oldPos := o.pos
	o.pos = nextPos
	objects[nextPos] = *o
	delete(objects, oldPos)

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
