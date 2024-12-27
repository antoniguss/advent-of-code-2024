package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const filename = "input.txt"

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

	//Seems like something we have to do manually, what we can do is add a useful visualisation

	zeros := make([]string, 0)
	xValues := make([]string, 0)
	yValues := make([]string, 0)
	xSum := 0
	ySum := 0
	for i := 0; ; i++ {
		z := "z" + fmt.Sprintf("%02d", i)
		x := "x" + fmt.Sprintf("%02d", i)
		y := "y" + fmt.Sprintf("%02d", i)
		if _, has := operations[z]; !has {
			break
		}
		zeros = append(zeros, z)
		xValues = append(xValues, x)
		xSum += values[x] << i
		yValues = append(yValues, y)
		ySum += values[y] << i
	}
	//
	//---Checking which bits give incorrect values
	swaps := []string{}
	//Found bit 6 wrong
	// FIXED: swap z06 and fhc
	operations["z06"], operations["fhc"] = operations["fhc"], operations["z06"]
	swaps = append(swaps, "z06", "fhc")
	//Found bit 11 wrong
	// FIXED: swap z10 and qhj
	operations["z11"], operations["qhj"] = operations["qhj"], operations["z11"]
	swaps = append(swaps, "z11", "qhj")

	//Found bit 23 wrong
	//FIXED: swap ggt and mwh
	operations["ggt"], operations["mwh"] = operations["mwh"], operations["ggt"]
	swaps = append(swaps, "ggt", "mwh")

	//Found bit 35 wrong
	//FIXED: swap z35 and hqk
	operations["z35"], operations["hqk"] = operations["hqk"], operations["z35"]
	swaps = append(swaps, "z35", "hqk")

	//After swap
	fmt.Printf("xSum: %v\n", xSum)
	fmt.Printf("ySum: %v\n", ySum)
	expectedSum := xSum + ySum
	fmt.Printf("expectedSum: %v\n", expectedSum)
	computer := computer{values: values, operations: operations}
	//i is the amout of first bits to check in input and output

	for i := 0; i < len(zeros); i++ {
		receivedBit := computer.compute(zeros[i])
		expectedBit := getBit(expectedSum, i)
		fmt.Printf("%d: calc: %d, exp: %d\n", i, receivedBit, expectedBit)
	}

	// --- visualizing the wires for every z wire
	var displayWire func(wire string, depth int)
	displayWire = func(wire string, depth int) {
		if _, has := values[wire]; has {
			fmt.Printf("%s%s\n", strings.Repeat(" ", depth), wire)
			return
		}

		op := operations[wire]
		fmt.Printf("%s%s (%s)\n", strings.Repeat(" ", depth), op.op, wire)
		displayWire(op.a, depth+1)
		displayWire(op.b, depth+1)
	}

	// for _, z := range zeros {
	// 	displayWire(z, 0)
	// 	util.Freeze()
	// }

	sort.Slice(swaps, func(i, j int) bool {
		return strings.Compare(swaps[i], swaps[j]) == -1
	})

	return strings.Join(swaps, ",")
}

func getBit(n int, i int) int {
	return (n >> i) & 1
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
		// fmt.Printf("%s: %s\n", a, op)
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
