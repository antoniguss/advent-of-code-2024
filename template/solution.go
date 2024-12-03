package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code - Day X") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line of input
		fmt.Println(line)
	}


}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
