package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type operation struct {
	values []int
	result int
}

// Check recursively
func checkOperation(input []int, result int) bool {
	if len(input) == 1 {
		return input[0] == result
	}

	inputAdd := make([]int, 1)
	inputMul := make([]int, 1)
	inputAdd[0] = input[0] + input[1]
	inputMul[0] = input[0] * input[1]

	inputAdd = append(inputAdd, input[2:]...)
	inputMul = append(inputMul, input[2:]...)
	return checkOperation(inputAdd, result) || checkOperation(inputMul, result)

}

func checkOperationWithConcat(input []int, result int) bool {
	if len(input) == 1 {
		return input[0] == result
	}
	if checkOperation(input, result) {
		return true
	}

	n1 := input[1]
	shift := 10
	n1 /= 10
	for n1 > 0 {
		shift *= 10
		n1 /= 10
	}

	inputAdd := make([]int, 1)
	inputMul := make([]int, 1)
	inputConcat := make([]int, 1)

	inputAdd[0] = input[0] + input[1]
	inputMul[0] = input[0] * input[1]
	inputConcat[0] = concatNumbers(input[0], input[1])

	inputAdd = append(inputAdd, input[2:]...)
	inputMul = append(inputMul, input[2:]...)
	inputConcat = append(inputConcat, input[2:]...)
	return checkOperationWithConcat(inputConcat, result) || checkOperationWithConcat(inputAdd, result) || checkOperationWithConcat(inputMul, result)

}

func concatNumbers(n1, n2 int) int {

	shift := 10
	temp := n2 / 10
	for temp > 0 {
		shift *= 10
		temp /= 10
	}

	return n1*shift + n2

}

func main() {
	fmt.Println("Advent of Code - Day 7")

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	scanner := bufio.NewScanner(file)
	operationRegex := regexp.MustCompile("([0-9]+):(( [0-9]+)+)")

	var operations []operation
	for scanner.Scan() {
		line := scanner.Text()

		submatch := operationRegex.FindStringSubmatch(line)
		result, _ := strconv.Atoi(submatch[1])

		valuesS := strings.Fields(submatch[2])
		values := make([]int, len(valuesS))
		for i, s := range valuesS {
			values[i], _ = strconv.Atoi(s)
		}

		operations = append(operations, operation{
			values: values,
			result: result,
		})
	}

	fmt.Println(operations)

	part1 := 0
	for _, op := range operations {
		if checkOperation(op.values, op.result) {
			part1 += op.result
		}
	}
	fmt.Printf("Part1: %d\n", part1)
	fmt.Println(len(operations))

	//--- Part 2 ---

	part2 := 0
	for _, op := range operations {
		if checkOperationWithConcat(op.values, op.result) {
			fmt.Printf("%v passed\n", op)
			part2 += op.result
		}
	}
	fmt.Printf("Part2: %d\n", part2)
	fmt.Println(concatNumbers(1, 2))
	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
