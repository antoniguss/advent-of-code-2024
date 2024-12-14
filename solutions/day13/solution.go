package main

import (
	"advent-of-code-2024/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type machine struct {
	A     util.Vector
	B     util.Vector
	Prize util.Vector
}

func main() {
	fmt.Println("Advent of Code - Day 13") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---

	buttonRegex := regexp.MustCompile("Button [A-Za-z]: X\\+([0-9]+), Y\\+([0-9]+)")
	prizeRegex := regexp.MustCompile("Prize: X=([0-9]+), Y=([0-9]+)")
	scanner := bufio.NewScanner(file)
	machines := make([]machine, 0)

	for scanner.Scan() {
		buttonALine := scanner.Text()
		Asub := buttonRegex.FindStringSubmatch(buttonALine)
		scanner.Scan()

		buttonBLine := scanner.Text()
		Bsub := buttonRegex.FindStringSubmatch(buttonBLine)
		scanner.Scan()

		prizeLine := scanner.Text()
		prizeSub := prizeRegex.FindStringSubmatch(prizeLine)
		scanner.Scan()

		Ax, _ := strconv.Atoi(Asub[1])
		Ay, _ := strconv.Atoi(Asub[2])

		Bx, _ := strconv.Atoi(Bsub[1])
		By, _ := strconv.Atoi(Bsub[2])

		prizex, _ := strconv.Atoi(prizeSub[1])
		prizey, _ := strconv.Atoi(prizeSub[2])

		newMachine := machine{
			A:     util.Vector{X: Ax, Y: Ay},
			B:     util.Vector{X: Bx, Y: By},
			Prize: util.Vector{X: prizex, Y: prizey},
		}
		machines = append(machines, newMachine)
	}

	tokens1 := part1(machines)
	fmt.Printf("Part1: %d\n", tokens1)

	//--- Part 2 ---
	tokens2 := part2(machines)
	fmt.Printf("Part2: %d\n", tokens2)

	//--- Cleanup ---
	err = file.Close()
	check(err)
}

func part1(machines []machine) (tokens int) {
	for _, machine := range machines {
		a1 := machine.A.X
		b1 := machine.B.X
		c1 := machine.Prize.X

		a2 := machine.A.Y
		b2 := machine.B.Y
		c2 := machine.Prize.Y

		//We're looking for x,y such that:
		//x*a1 +y*b1 = c1
		//x*a2 +y*b2 = c2

		detA := a1*b2 - a2*b1

		if detA == 0 {
			//Solution doesn't exist
			continue
		}

		detA1 := c1*b2 - b1*c2
		detA2 := a1*c2 - a2*c1

		//x = detA1/detA, but x must be natural so detA | detA1
		if detA1%detA != 0 {
			continue
		}

		if detA2%detA != 0 {
			continue
		}

		x := detA1 / detA
		y := detA2 / detA

		tokens += x*3 + y

	}

	return tokens
}

func part2(machines []machine) (tokens int) {
	newMachines := make([]machine, len(machines))
	for i, m := range machines {
		newMachine := machine{
			A:     m.A,
			B:     m.B,
			Prize: m.Prize.Add(util.Vector{X: 10000000000000, Y: 10000000000000}),
		}
		newMachines[i] = newMachine

	}

	return part1(newMachines) // Placeholder for part 2 logic
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
