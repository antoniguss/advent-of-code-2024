package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "input.txt"

func main() {
	fmt.Println("Advent of Code - Day 23") // Placeholder for day number

	connections, computers := getInput()

	// for c1, conns := range connections {
	//
	// 	fmt.Printf("%s:", c1)
	// 	for c2 := range conns {
	// 		fmt.Printf(" %s", c2)
	// 	}
	// 	fmt.Println()
	//
	// }

	res1 := part1(connections, computers)
	fmt.Printf("Part1: %d\n", res1)

}

func part1(connections map[string]map[string]struct{}, computers map[string]struct{}) (result int) {

	for x := range connections {
		for y := range connections[x] {
			for z := range connections[y] {

				if _, has := connections[z][x]; has {
					if x[0] == 't' || y[0] == 't' || z[0] == 't' {
						result++

					}
				}

			}

		}

	}

	return result / 6
}

func getInput() (connections map[string](map[string]struct{}), computers map[string]struct{}) {
	file, err := os.Open(filename)
	check(err)

	connections = make(map[string]map[string]struct{})
	computers = make(map[string]struct{})
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "-")

		if _, has := connections[c[0]]; !has {
			connections[c[0]] = make(map[string]struct{})
		}

		if _, has := connections[c[1]]; !has {
			connections[c[1]] = make(map[string]struct{})
		}

		connections[c[0]][c[1]] = struct{}{}
		connections[c[1]][c[0]] = struct{}{}

	}

	return connections, computers

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
