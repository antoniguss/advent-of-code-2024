package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type update []int

func main() {
	fmt.Println("Advent of Code - Day 5") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	//--- Part 1 ---
	//After some code cleanup Part1 and Part2 are computed at the same time

	//Process rules
	//Here the second map is used as a set of elements that should appear after the key
	before := make(map[int]map[int]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, "|")
		n0, err := strconv.Atoi(split[0])
		check(err)
		n1, err := strconv.Atoi(split[1])
		check(err)

		if _, has := before[n0]; !has {
			before[n0] = make(map[int]struct{})
		}
		before[n0][n1] = struct{}{}
	}

	//Process updates
	var updates []update
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		u := make(update, len(split))

		for i, s := range split {
			n, err := strconv.Atoi(s)
			check(err)
			u[i] = n
		}
		updates = append(updates, u)
	}

	sum1 := 0
	sum2 := 0
	for _, u := range updates {
		if !fixUpdate(u, before) {
			sum2 += u[len(u)/2]
		} else {
			sum1 += u[len(u)/2]
		}

	}
	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func fixUpdate(u update, before map[int]map[int]struct{}) bool {
	correct := true
	for i := len(u) - 1; i > 0; {
		swapped := false
		for i2 := 0; i2 < i; i2++ {
			if _, has := before[u[i]][u[i2]]; has {
				correct = false
				u[i], u[i2] = u[i2], u[i]
				swapped = true
				break
			}
		}
		if !swapped {
			i--
		}
	}
	return correct
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
