package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "input.txt"

func main() {
	fmt.Println("Advent of Code - Day 22") // Placeholder for day number

	numbers := getInput()

	res1 := part1(numbers)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(numbers)
	fmt.Printf("Part2: %d\n", res2)

}

func part1(numbers []int) (result int) {
	iter := 2000

	for _, num := range numbers {
		// fmt.Printf("num: %v\n", num)
		secret := num
		for range iter {
			// Mul by 64 and mix
			n := secret << 6
			secret = n ^ secret
			secret = secret & ((1 << 24) - 1)

			// Div by 32, mix and prune
			n = secret >> 5
			secret = n ^ secret
			secret = secret & ((1 << 24) - 1)

			// Mul by 2048, mix and prune
			n = secret << 11
			secret = n ^ secret
			secret = secret & ((1 << 24) - 1)

		}
		fmt.Println(num, secret)
		result += secret
	}

	return result
}

func part2(numbers []int) (result int) {
	return result
}

func getInput() (numbers []int) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
