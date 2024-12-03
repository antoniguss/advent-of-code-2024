package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fileName := "./input.txt"
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)

	// --- Part 1 ---
	var left []int
	var right []int

	occurrences := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		n1, err := strconv.Atoi(fields[0])
		check(err)
		n2, err := strconv.Atoi(fields[1])
		check(err)
		left = append(left, n1)
		right = append(right, n2)

		//Preprocessing data for part 2
		occurrences[n2]++

	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] > left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] > right[j]
	})

	sum := 0
	for i := range left {
		sum += abs(left[i] - right[i])
	}

	fmt.Printf("Final sum: %d\n", sum)
	// --- Part 2 ---
	//  Calculate a total similarity score by adding up each number in the left list
	//  after multiplying it by the number of times that number appears in the right
	//  list.

	score := 0
	for _, val := range left {
		score += occurrences[val] * val
	}

	fmt.Printf("Final score: %d\n", score)

	err = f.Close()
	check(err)

}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
