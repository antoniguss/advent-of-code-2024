package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const filename = "input.txt"

func main() {
	fmt.Println("Advent of Code - Day 24") // Placeholder for day number

	values, operations := getInput()

	start := time.Now()
	res1 := part1(values, operations)
	fmt.Printf("Part1: %v, took %v\n", res1, time.Since(start))

	start = time.Now()
	res2 := part2()
	fmt.Printf("Part2: %v, took %v\n", res2, time.Since(start))

}

func part1(values map[string]int, operations map[string]operation) (result int) {
	zeros := make([]string, 0)
	for i := 0; ; i++ {
		z := "z" + fmt.Sprintf("%02d", i)
		if _, has := operations[z]; !has {
			break
		}
		zeros = append(zeros, z)
	}
	computer := computer{values: values, operations: operations}

	for i, z := range zeros {
		val := computer.compute(z)
		// fmt.Printf("%s: %d\n", z, val)
		result += val << i
	}

	return result
}

func (c computer) compute(a string) int {
	if val, has := c.values[a]; has {
		return val
	}

	if op, has := c.operations[a]; has {
		switch op.op {
		case "AND":
			{
				return c.compute(op.a) & c.compute(op.b)

			}
		case "OR":
			{
				return c.compute(op.a) | c.compute(op.b)

			}
		case "XOR":
			{
				return c.compute(op.a) ^ c.compute(op.b)

			}
		}

	}

	panic("COULDN'T compute " + a)

}

type computer struct {
	values     map[string]int
	operations map[string]operation
}

func part2() (result int) {
	return result
}

func getInput() (map[string]int, map[string]operation) {

	file, err := os.Open(filename)
	check(err)

	values := make(map[string]int)
	operations := make(map[string]operation)

	//--- Part 1 ---
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, ": ")
		values[split[0]], _ = strconv.Atoi(split[1])
	}

	for scanner.Scan() {
		line := scanner.Text()
		splitArr := strings.Split(line, " -> ")
		splitOp := strings.Fields(splitArr[0])
		operations[splitArr[1]] = operation{a: splitOp[0], op: splitOp[1], b: splitOp[2]}
	}

	return values, operations
}

type operation struct {
	a  string
	op string
	b  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
