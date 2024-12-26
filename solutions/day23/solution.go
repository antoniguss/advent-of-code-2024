package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const filename = "input.txt"

func main() {
	fmt.Println("Advent of Code - Day 23") // Placeholder for day number

	connections := getInput()

	// for c1, conns := range connections {
	//
	// 	fmt.Printf("%s:", c1)
	// 	for c2 := range conns {
	// 		fmt.Printf(" %s", c2)
	// 	}
	// 	fmt.Println()
	//
	// }

	res1 := part1(connections)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(connections)
	fmt.Printf("Part2: %s\n", res2)
}

func part1(connections map[string]map[string]struct{}) (result int) {

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

func setToString(set map[string]struct{}) (output string) {

	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) == -1
	})

	return strings.Join(keys, ",")

}

func part2(connections map[string]map[string]struct{}) (result string) {

	sets := make(map[string]struct{})

	var searchFunc func(node string, req map[string]struct{})
	searchFunc = func(node string, req map[string]struct{}) {
		key := setToString(req)
		if _, has := sets[key]; has {
			return
		}

		sets[key] = struct{}{}

		for neighbour := range connections[node] {
			if _, has := req[neighbour]; has {
				continue
			}

			// We make sure this neighbour connects to all required nodes
			if len(connections[neighbour]) < len(req) {
				continue
			}
			notAll := false
			for query := range req {
				if _, has := connections[neighbour][query]; !has {
					notAll = true
					break
				}
			}
			if notAll {
				continue
			}

			req[neighbour] = struct{}{}

			searchFunc(neighbour, req)

		}
	}

	for node := range connections {
		newReq := map[string]struct{}{node: {}}
		searchFunc(node, newReq)
	}

	maxLan := ""
	for lan := range sets {
		if len(lan) > len(maxLan) {
			maxLan = lan
		}

	}

	return maxLan
}

func getInput() (connections map[string](map[string]struct{})) {
	file, err := os.Open(filename)
	check(err)

	connections = make(map[string]map[string]struct{})
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

	return connections

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
