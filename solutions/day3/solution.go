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

	//IDEA: Use regex, something like 'mul(%d,%d)' to accept a mul command
	regex, err := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	check(err)

	allString := regex.FindAllString(inputString, -1)
	sum := 0
	for _, operation := range allString {
		regex = regexp.MustCompile("\\((\\d+),(\\d+)\\)")
		submatch := regex.FindStringSubmatch(operation)
		n1, err := strconv.Atoi(submatch[1])
		check(err)
		n2, err := strconv.Atoi(submatch[2])
		check(err)

		sum += n1 * n2
	}

	fmt.Printf("The sum of all operations is: %d\n", sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
