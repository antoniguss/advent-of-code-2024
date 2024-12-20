package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day 18") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line of input
		fmt.Println(line)
	}

	//--- Part 2 ---

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
