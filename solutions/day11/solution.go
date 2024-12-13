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
	sum := part1(inputList, 25)
	fmt.Printf("Part 1: %d\n", sum)
	//--- Part 2 ---

	sum2 := part2(inputList, 50)
	fmt.Printf("Part 2: %d\n", sum2)
	//--- Cleanup ---
	err = file.Close()
	check(err)

}

func part2(inputList util.LinkedList, iterations int) int {
	//IDEA: Use map to store count of differnt values
	curr := inputList.Head
	counts := make(map[int]int)
	for curr != nil {
		counts[curr.Value]++
		curr = curr.Next
	}

	for range iterations {
		newMap := make(map[int]int)

		for value, count := range counts {

			length := intLength(value)
			if value == 0 {
				newMap[1] += count
			} else if length%2 == 0 {
				left := value / int(math.Pow10(length/2))
				right := value - left*int(math.Pow10(length/2))

				newMap[left] += count
				newMap[right] += count
			} else {
				newMap[value*2024] += count
			}
		}
		counts = newMap

	}

	sum := 0
	for _, count := range counts {
		sum += count
	}

	return sum

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

//------------Thought recursion migth be smart but no :(--------------------
// func part2(inputList util.LinkedList, maxDepth int) int {
// 	curr := inputList.Head length := 0 for curr != nil {
// 		length++
// 		curr = curr.Next
// 	}
//
// 	sum := 0
// 	curr = inputList.Head
// 	count := 0
// 	for curr != nil {
// 		sum += recurse(curr.Value, maxDepth)
// 		count++
// 		fmt.Printf("(%d/%d)\n", count, length)
//
// 		curr = curr.Next
// 	}
//
// 	return sum
// }

// func recurse(value int, remainDepth int) int {
// 	if remainDepth == 0 {
// 		return 1
// 	}
//
// 	length := intLength(value)
// 	if value == 0 {
// 		return recurse(1, remainDepth-1)
// 	} else if length%2 == 0 {
// 		left := value / int(math.Pow10(length/2))
// 		right := value - left*int(math.Pow10(length/2))
//
// 		return recurse(left, remainDepth-1) + recurse(right, remainDepth-1)
// 	} else {
// 		return recurse(value*2024, remainDepth-1)
// 	}
// }
