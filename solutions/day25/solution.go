package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "input.txt"

func main() {
	fmt.Println("Advent of Code - Day 25") // Placeholder for day number

	locks, keys := getInput()

	sol1 := part1(locks, keys)
	fmt.Printf("Part1: %d\n", sol1)

	sol2 := part2(locks, keys)
	fmt.Printf("Part2: %d\n", sol2)
}

func part1(locks, keys []schematic) (result int) {
	for _, lock := range locks {
		for _, key := range keys {
			if lock.fits(key) {
				result++
			}

		}
	}

	return result
}

func part2(locks, keys []schematic) (result int) {
	return result
}

type schematic [5]int

func (lock schematic) fits(key schematic) bool {
	space := 5
	for i := range 5 {
		if lock[i]+key[i] > space {
			return false
		}

	}

	return true

}

func getInput() (locks, keys []schematic) {

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)

	for {
		if !scanner.Scan() {
			break
		}

		isLock := scanner.Text()[0] == '#'
		s := schematic{}

		for line := scanner.Text(); len(line) > 0; line = scanner.Text() {

			for i, c := range line {
				if c == '#' {
					s[i]++
				}
			}
			scanner.Scan()

		}

		for i := range 5 {
			s[i]--
		}

		if isLock {
			locks = append(locks, s)
		} else {
			keys = append(keys, s)
		}
	}

	return locks, keys
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
