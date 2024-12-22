package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(inputPath string) (towels []string, wanted map[string]struct{}, longest int) {
	file, err := os.Open(inputPath)
	check(err)
	wanted = make(map[string]struct{})

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	temp := strings.Split(line, ", ")
	for _, towel := range temp {
		towels = append(towels, strings.TrimSpace(towel))
	}

	//Load wanted paterns
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		// Process each line of input
		wanted[line] = struct{}{}
		if len(line) > longest {
			longest = len(line)
		}
	}
	return towels, wanted, longest
}

func main() {
	fmt.Println("Advent of Code - Day 19") // Placeholder for day number

	inputPath := "./input.txt"
	towels, wanted, longest := input(inputPath)

	// fmt.Println(towels, wanted)

	//--- Part 1 ---
	sol1 := part1(towels, wanted, longest)
	fmt.Printf("Part1: %d\n", sol1)

	//--- Part 2 ---
	// sol2 := part2(towels, wanted)
	// fmt.Printf("Part2: %d\n", sol2)

}

func part2(towels, wanted []string) (count int) {

	return count
}

func part1(towels []string, wanted map[string]struct{}, longest int) (count int) {
	made := make(map[string]struct{})
	notMade := make(map[string]struct{})

	for pattern := range wanted {
		// fmt.Printf("pattern: %v\n", pattern)
		if canMake(towels, pattern, made, notMade) {
			count++
		}

	}
	return count
}

func canMake(towels []string, pattern string, made, notMade map[string]struct{}) bool {
	if len(pattern) == 0 {
		return true
	}

	if _, has := made[pattern]; has {
		return true
	}
	if _, has := notMade[pattern]; has {
		return false
	}
	// fmt.Println(len(made))

	for _, towel := range towels {
		// fmt.Printf("towel: %v\n", towel)
		if len(pattern) >= len(towel) && towel == pattern[:len(towel)] {
			if canMake(towels, pattern[len(towel):], made, notMade) {
				made[pattern[len(towel):]] = struct{}{}
				return true
			}
		}

	}
	// fmt.Println("Coulnd't")

	notMade[pattern] = struct{}{}
	return false

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
