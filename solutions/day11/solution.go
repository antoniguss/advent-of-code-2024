package main

import (
	"advent-of-code-2024/util"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 11") // Placeholder for day number

	file, err := os.Open("input.txt")
	check(err)

	//--- Part 1 ---
	inputList := util.LinkedList{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	values := strings.Fields(line)
	fmt.Println(values)

	var current *util.Node = inputList.Head
	for _, s := range values {
		val, _ := strconv.Atoi(s)

		if current == nil {
			current = &util.Node{Value: val}
			inputList.Head = current
		} else {
			newNode := &util.Node{Value: val}
			current.Next = newNode
			current = newNode
		}

	}
	inputList.Print()
	sum := part1(inputList, 75)
	fmt.Printf("Part 1: %d\n", sum)
	//--- Part 2 ---

	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part1(inputList util.LinkedList, iterations int) int {

	for i := range iterations {
		fmt.Println(i)
		var curr *util.Node = inputList.Head

		for curr != nil {

			//If number 0, replace with 1
			length := intLength(curr.Value)
			if curr.Value == 0 {
				curr.Value = 1
				curr = curr.Next
			} else if length%2 == 0 {
				nextNode := curr.Next
				left := curr.Value / int(math.Pow10(length/2))
				right := curr.Value - left*int(math.Pow10(length/2))

				curr.Value = left
				newNode := &util.Node{Value: right, Next: nextNode}
				curr.Next = newNode
				curr = nextNode
			} else {
				//If even number of digits, replace by two stones (leftSide) (rightSide)

				//If none, multiply value by 2024
				curr.Value = curr.Value * 2024
				curr = curr.Next
			}
		}
	}

	curr := inputList.Head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}

	return count
}

func intLength(n int) int {
	length := 1
	n /= 10
	for n > 0 {
		length++
		n /= 10
	}

	return length
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
