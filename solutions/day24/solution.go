package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const filename = "test2.txt"

func main() {
	fmt.Println("Advent of Code - Day 24") // Placeholder for day number

	x, y, values, operations := getInput()

	start := time.Now()
	res1 := part1(values, operations)
	fmt.Printf("Part1: %v, took %v\n", res1, time.Since(start))

	start = time.Now()
	res2 := part2(x, y, values, operations)
	fmt.Printf("Part2: %v, took %v\n", res2, time.Since(start))

}

func part2(x, y int, values map[string]int, operations map[string]operation) (result string) {

	// Go through all possible 4 pairs to swap
	// We won't track the ones we tested as finding the first set of pairs will give us the result

	zeros := make([]string, 0)
	for i := 0; ; i++ {
		z := "z" + fmt.Sprintf("%02d", i)
		if _, has := operations[z]; !has {
			break
		}
		zeros = append(zeros, z)
	}

	for _, z := range zeros {
		fmt.Printf("%s: %v\n", z, operations[z])
	}

	return ""
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

func (c computer) compute(a string) (result int) {

	if val, has := c.values[a]; has {
		// fmt.Printf("%s: %d\n", a, val)
		return val
	}

	if op, has := c.operations[a]; has {
		fmt.Printf("%s: %s\n", a, op)
		switch op.op {
		case "AND":
			{
				result = c.compute(op.a) & c.compute(op.b)

			}
		case "OR":
			{
				result = c.compute(op.a) | c.compute(op.b)

			}
		case "XOR":
			{
				result = c.compute(op.a) ^ c.compute(op.b)

			}
		}

	}

	//Enable to "cache" results, doesn't change efficiency of part1, will see for part2
	// c.values[a] = result
	return result

}

type computer struct {
	values     map[string]int
	operations map[string]operation
}

func getInput() (int, int, map[string]int, map[string]operation) {

	file, err := os.Open(filename)
	check(err)

	values := make(map[string]int)
	operations := make(map[string]operation)
	i := 0
	j := 0
	x := 0
	y := 0

	//--- Part 1 ---
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, ": ")
		val, _ := strconv.Atoi(split[1])
		values[split[0]] = val
		if split[0][0] == 'x' {
			x += val << i
			i++
		}
		if split[0][0] == 'y' {
			y += val << j
			j++
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		splitArr := strings.Split(line, " -> ")
		splitOp := strings.Fields(splitArr[0])
		operations[splitArr[1]] = operation{a: splitOp[0], op: splitOp[1], b: splitOp[2]}
	}

	return x, y, values, operations
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
