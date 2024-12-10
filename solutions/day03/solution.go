package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code - Day 3") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputString := ""
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line of input
		inputString += line
	}

	//IDEA: Use mulRegex, something like 'mul(%d,%d)' to accept a mul command
	mulRegex, err := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	check(err)

	allString := mulRegex.FindAllString(inputString, -1)
	sum := 0
	for _, operation := range allString {
		sum += processMul(operation)
	}

	fmt.Printf("The sum of all operations is: %d\n", sum)

	//--- Part 2 ---
	sum2 := 0
	doRegex := regexp.MustCompile("do\\(\\)")
	dontRegex := regexp.MustCompile("don't\\(\\)")
	doIndices := doRegex.FindAllStringIndex(inputString, -1)
	dontIndices := dontRegex.FindAllStringIndex(inputString, -1)
	mulIndicies := mulRegex.FindAllStringIndex(inputString, -1)

	fmt.Println(doIndices)
	fmt.Println(dontIndices)
	fmt.Println(mulIndicies)
	do := true
	for i := range len(inputString) {
		if len(doIndices) > 0 && doIndices[0][0] == i {
			do = true
			doIndices = append(doIndices[:0], doIndices[1:]...)
		}

		if len(dontIndices) > 0 && dontIndices[0][0] == i {
			do = false
			dontIndices = append(dontIndices[:0], dontIndices[1:]...)
		}

		if len(mulIndicies) > 0 && mulIndicies[0][0] == i {
			mulIndex := mulIndicies[0]
			if do {
				mulString := inputString[mulIndex[0]:mulIndex[1]]
				sum2 += processMul(mulString)
			}

			mulIndicies = append(mulIndicies[:0], mulIndicies[1:]...)

		}
	}

	fmt.Printf("The sum of all operations is: %d\n", sum2)

}

func processMul(operation string) int {
	parameterRegex := regexp.MustCompile("\\((\\d+),(\\d+)\\)")
	submatch := parameterRegex.FindStringSubmatch(operation)
	n1, err := strconv.Atoi(submatch[1])
	check(err)
	n2, err := strconv.Atoi(submatch[2])
	check(err)
	return n1 * n2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
