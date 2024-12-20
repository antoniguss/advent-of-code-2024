package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type state struct {
	regA         int
	regB         int
	regC         int
	instructions []int
	ip           int
}

func main() {
	fmt.Println("Advent of Code - Day 17") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	var regA, regB, regC int

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	regA, _ = strconv.Atoi(strings.Fields(scanner.Text())[2])

	scanner.Scan()
	regB, _ = strconv.Atoi(strings.Fields(scanner.Text())[2])

	scanner.Scan()
	regC, _ = strconv.Atoi(strings.Fields(scanner.Text())[2])

	fmt.Println(regA, regB, regC)

	scanner.Scan()
	scanner.Scan()
	temp := strings.Split(strings.Split(scanner.Text(), " ")[1], ",")
	instructions := make([]int, 0)
	for _, val := range temp {
		instruction, _ := strconv.Atoi(val)
		instructions = append(instructions, instruction)
	}

	state1 := state{
		regA:         regA,
		regB:         regB,
		regC:         regC,
		instructions: instructions,
	}
	fmt.Printf("instructions: %v\n", instructions)

	output1 := state1.part1()
	fmt.Printf("Part1: %s\n", output1)

	//--- Part 2 --
	output2 := part2(regA, state1.instructions)
	fmt.Printf("Part2: %d\n", output2)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func (s *state) combo(operand int) int {
	switch {
	case operand <= 3:
		{
			return operand
		}
	case operand == 4:
		{
			return s.regA
		}
	case operand == 5:
		{
			return s.regB
		}
	case operand == 6:
		{
			return s.regC
		}
	}

	return -1

}

func (s *state) part1() (outputs []int) {

	for s.ip < len(s.instructions) {
		opcode := s.instructions[s.ip]
		operand := s.instructions[s.ip+1]
		// fmt.Println(s)
		// fmt.Println(opcode, operand)

		switch opcode {
		case 0:
			{
				// fmt.Println("adv")
				s.regA = s.regA >> s.combo(operand)
				// fmt.Println(s.regA)

				s.ip += 2
			}
		case 1:
			{
				// old := s.regB
				s.regB = s.regB ^ operand
				// fmt.Println("bxl")
				// fmt.Printf("%d xor %d = %d\n", old, operand, s.regB)

				s.ip += 2
			}
		case 2:
			{
				// fmt.Println("bst")
				s.regB = s.combo(operand) & 7
				// fmt.Printf("%d mod 8 = %d\n", s.combo(operand), s.regB)

				s.ip += 2

			}
		case 3:
			{
				// fmt.Println("jnz")
				if s.regA == 0 {
					s.ip += 2
				} else {
					s.ip = operand
					// fmt.Printf("ip = %d\n", s.ip)
				}
			}
		case 4:
			{

				// fmt.Println("bxc")
				s.regB = s.regB ^ s.regC

				s.ip += 2
			}
		case 5:
			{

				// fmt.Println("out")
				outputs = append(outputs, s.combo(operand)&7)
				s.ip += 2
			}
		case 6:
			{

				// fmt.Println("bvd")
				numerator := s.regA
				denumerator := int(math.Pow(2, float64(s.combo(operand))))
				// fmt.Printf("%d/%d = %d\n", numerator, denumerator, numerator/denumerator)
				s.regB = numerator / denumerator

				s.ip += 2
			}
		case 7:
			{

				// fmt.Println("cvd")
				numerator := s.regA
				denumerator := int(math.Pow(2, float64(s.combo(operand))))
				// fmt.Printf("%d/%d = %d\n", numerator, denumerator, numerator/denumerator)
				s.regC = numerator / denumerator
				s.ip += 2
			}
		}

	}

	// fmt.Printf("outputs: %v\n", outputs)
	return outputs
}

func part2(prevA int, instructions []int) (value int) {
	//Brutforcing probably not worth it, will do manually
	solution, _ := find(instructions, 0)

	return solution
}

func find(instructions []int, ans int) (sub int, found bool) {

	if len(instructions) == 0 {
		return ans, true
	}

	for i := 0; i < 8; i++ {
		b := i
		a := ans<<3 + b

		b = b ^ 2
		c := a >> b
		b = b ^ c
		b = b ^ 3
		if b%8 == instructions[len(instructions)-1] {
			sub, found := find(instructions[:len(instructions)-1], a)
			if !found {
				continue
			}
			return sub, true
		}

	}
	return -1, false

}
func (s *state) compare(expected []int) bool {

	i := 0
	for s.ip < len(s.instructions) {
		if i >= len(expected) {
			return false
		}
		opcode := s.instructions[s.ip]
		operand := s.instructions[s.ip+1]
		// fmt.Println(s)
		// fmt.Println(opcode, operand)

		switch opcode {
		case 0:
			{
				// fmt.Println("adv")
				s.regA = s.regA >> s.combo(operand)
				// fmt.Println(s.regA)

				s.ip += 2
			}
		case 1:
			{
				// old := s.regB
				s.regB = s.regB ^ operand
				// fmt.Println("bxl")
				// fmt.Printf("%d xor %d = %d\n", old, operand, s.regB)

				s.ip += 2
			}
		case 2:
			{
				// fmt.Println("bst")
				s.regB = s.combo(operand) & 7
				// fmt.Printf("%d mod 8 = %d\n", s.combo(operand), s.regB)

				s.ip += 2

			}
		case 3:
			{
				// fmt.Println("jnz")
				if s.regA == 0 {
					s.ip += 2
				} else {
					s.ip = operand
					// fmt.Printf("ip = %d\n", s.ip)
				}
			}
		case 4:
			{

				// fmt.Println("bxc")
				s.regB = s.regB ^ s.regC

				s.ip += 2
			}
		case 5:
			{

				// fmt.Println("out")
				out := s.combo(operand) & 7
				if expected[i] != out {
					return false
				}
				s.ip += 2
				i++
			}
		case 6:
			{

				// fmt.Println("bvd")
				numerator := s.regA
				denumerator := int(math.Pow(2, float64(s.combo(operand))))
				// fmt.Printf("%d/%d = %d\n", numerator, denumerator, numerator/denumerator)
				s.regB = numerator / denumerator

				s.ip += 2
			}
		case 7:
			{

				// fmt.Println("cvd")
				numerator := s.regA
				denumerator := int(math.Pow(2, float64(s.combo(operand))))
				// fmt.Printf("%d/%d = %d\n", numerator, denumerator, numerator/denumerator)
				s.regC = numerator / denumerator
				s.ip += 2
			}
		}

	}

	// fmt.Printf("outputs: %v\n", outputs)
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
