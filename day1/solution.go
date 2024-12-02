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
	fmt.Println("Hello World!")

	fileName := "./input.txt"
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)

	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		n1, err := strconv.Atoi(fields[0])
		check(err)
		n2, err := strconv.Atoi(fields[1])
		check(err)
		left = append(left, n1)
		right = append(right, n2)

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
