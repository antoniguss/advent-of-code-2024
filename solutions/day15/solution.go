package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Antyhot/advent-of-code-24/util"
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
	objects2 := make(map[util.Vector]object, 0)
	var robot object
	var robot2 object

	for len(line) != 0 && line[0] == '#' {
		// Process each line of input
		for col, c := range line {
			pos := util.Vector{X: col, Y: rows}

			pos2 := util.Vector{X: col * 2, Y: rows}
			switch c {
			case '#':
				objects[pos] = object{pos: pos, isMovable: false, isRobot: false}

				object2 := object{pos: pos2, isMovable: false, isRobot: false}
				objects2[pos2] = object2
				objects2[pos2.Add(util.Vector{X: 1, Y: 0})] = object2

			case 'O':
				objects[pos] = object{pos: pos, isMovable: true, isRobot: false}
				object2 := object{pos: pos2, isMovable: true, isRobot: false}
				objects2[pos2] = object2
				objects2[pos2.Add(util.Vector{X: 1, Y: 0})] = object2
			case '@':
				robot = object{pos: pos, isMovable: true, isRobot: true}
				robot2 = object{pos: pos2, isMovable: true, isRobot: true}
				objects[pos] = robot
				objects2[pos2] = robot2
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
	sum2 := part2(objects2, instructions, robot2, rows, cols*2)
	fmt.Printf("Part2: %d\n", sum2)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(objects map[util.Vector]object, instruction []util.Vector, robot object, rows, cols int) (sum int) {
	// fmt.Println("Initial state:")
	//printObjects(objects, rows, cols)
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
func part2(objects map[util.Vector]object, instruction []util.Vector, robot object, rows, cols int) (sum int) {
	// fmt.Println("Initial state:")
	//printObjects(objects, rows, cols)
	for _, instruction := range instruction {
		//switch {
		//case instruction.X == 0 && instruction.Y == -1:
		//	fmt.Println("Move ^")
		//case instruction.X == 1 && instruction.Y == 0:
		//	fmt.Println("Move >")
		//case instruction.X == 0 && instruction.Y == 1:
		//	fmt.Println("Move v")
		//case instruction.X == -1 && instruction.Y == 0:
		//	fmt.Println("Move <")
		//}
		robot.move2(objects, instruction)
		//printObjects(objects, rows, cols)

	}

	checked := map[object]struct{}{}
	for _, object := range objects {

		if object.isMovable && !object.isRobot {
			if _, has := checked[object]; !has {
				sum += 100*object.pos.Y + object.pos.X
				checked[object] = struct{}{}
			}
		}
	}
	return sum
}

func printObjects(objects map[util.Vector]object, rows, cols int) {
	for row := range rows {
		for col := range cols {
			pos := util.Vector{X: col, Y: row}
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

func canMove(position util.Vector, objects map[util.Vector]object, instruction util.Vector) bool {

	o, has := objects[position]
	if !has {
		return true
	}

	if !o.isMovable {
		return false
	}

	nextPos := o.pos.Add(instruction)
	//If move is horizontal
	if instruction.Y == 0 {

		if o.isRobot {
			return canMove(nextPos, objects, instruction)
		}

		if instruction.X == 1 {
			nextPos = nextPos.Add(instruction)
		}

		return canMove(nextPos, objects, instruction)
	}
	//If vertical
	if o.isRobot {
		return canMove(nextPos, objects, instruction)
	}

	nextPosRight := nextPos.Add(util.Vector{X: 1, Y: 0})
	return canMove(nextPos, objects, instruction) && canMove(nextPosRight, objects, instruction)
}

func (o *object) move2(objects map[util.Vector]object, instruction util.Vector) {

	if !o.isMovable || !canMove(o.pos, objects, instruction) {
		fmt.Println("couldn't move")
		return
	}
	nextPos := o.pos.Add(instruction)

	if o.isRobot {
		if inFront, has := objects[nextPos]; has {
			inFront.move2(objects, instruction)
		}

		oldPos := o.pos
		o.pos = nextPos
		objects[o.pos] = *o
		delete(objects, oldPos)
		return

	}

	//If horizontal
	if instruction.Y == 0 {
		if instruction.X == -1 {
			if inFront, has := objects[nextPos]; has {
				inFront.move2(objects, instruction)
			}

			oldPos := o.pos
			o.pos = nextPos
			objects[o.pos] = *o
			objects[oldPos] = *o
			delete(objects, oldPos.Add(util.Vector{X: 1, Y: 0}))
			return

		} else {
			nextNextPos := nextPos.Add(instruction)
			if inFront, has := objects[nextNextPos]; has {
				inFront.move2(objects, instruction)
			}

			oldPos := o.pos
			o.pos = nextPos
			objects[o.pos] = *o
			objects[nextNextPos] = *o
			delete(objects, oldPos)
			return
		}
	}

	nextPosRight := nextPos.Add(util.Vector{X: 1, Y: 0})
	inFront, has := objects[nextPos]
	inFrontRight, hasRight := objects[nextPosRight]

	if has && hasRight {
		if inFront == inFrontRight {
			inFront.move2(objects, instruction)
		} else {
			inFront.move2(objects, instruction)
			inFrontRight.move2(objects, instruction)
		}
	} else {

		if has {
			inFront.move2(objects, instruction)
		}

		if hasRight {
			inFrontRight.move2(objects, instruction)
		}
	}

	oldPos := o.pos
	o.pos = nextPos
	objects[nextPos] = *o
	objects[nextPosRight] = *o
	delete(objects, oldPos)
	delete(objects, oldPos.Add(util.Vector{X: 1, Y: 0}))

	return

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
