package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "input.txt"

func main() {
	fmt.Println("Advent of Code - Day 22") // Placeholder for day number

	numbers := getInput()

	res1, sequences := part1(numbers)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(sequences)
	fmt.Printf("Part2: %d\n", res2)

}

func part1(numbers []int) (result int, sequences [][]int) {
	iter := 2000

	for _, num := range numbers {
		sequence := make([]int, iter)
		// fmt.Printf("num: %v\n", num)
		secret := num
		sequence[0] = secret
		for i := 1; i < iter; i++ {
			// Mul by 64 and mix
			n := secret << 6
			secret = n ^ secret
			secret = secret & ((1 << 24) - 1)

			// Div by 32, mix and prune
			n = secret >> 5
			secret = n ^ secret
			secret = secret & ((1 << 24) - 1)

			// Mul by 2048, mix and prune
			n = secret << 11
			secret = n ^ secret
			secret = secret & ((1 << 24) - 1)
			sequence[i] = secret

		}
		// fmt.Println(num, secret)
		result += secret
		sequences = append(sequences, sequence)
	}

	return result, sequences
}

func part2(sequences [][]int) (result int) {
	allChanges := make([][]int, len(sequences))

	for i, sequence := range sequences {
		// fmt.Println(sequence)

		changes := make([]int, len(sequences[0]))

		for j := 1; j < len(sequence)-1; j++ {
			change := sequence[j]%10 - sequence[j-1]%10
			changes[j] = change
		}

		// for i := 0; i < len(sequence); i++ {
		// 	fmt.Printf("%d: %d (%d)\n", sequence[i], sequence[i]%10, changes[i])
		// }
		allChanges[i] = changes
	}

	priceChangesToBananas := make(map[[4]int][]int)
	for i, sequence := range sequences {
		changes := allChanges[i]
		for j := 1; j < len(sequence)-4; j++ {
			priceChanges := [4]int{changes[j], changes[j+1], changes[j+2], changes[j+3]}

			if _, has := priceChangesToBananas[priceChanges]; !has {
				priceChangesToBananas[priceChanges] = make([]int, len(sequences))
			}
			if priceChangesToBananas[priceChanges][i] == 0 {
				priceChangesToBananas[priceChanges][i] = sequence[j+3] % 10

			}
		}
	}

	var maxSum int
	for _, bananas := range priceChangesToBananas {
		sum := 0
		for _, n := range bananas {
			sum += n
		}

		if sum > maxSum {
			maxSum = sum
		}

	}
	return maxSum
}

func getInput() (numbers []int) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
